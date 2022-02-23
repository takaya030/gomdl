package studio

import (
	"unsafe"
	"github.com/go-gl/gl/v2.1/gl"
)

var g_texnum int32 = 100

// skin info
type Texture struct {
	Name   [64]byte
	Flags  int32
	Width  int32
	Height int32
	Index  int32
}

type Rgb struct {
	r, g, b byte
}

func (tx *Texture) GetPixelBuf(basebuf []byte) []byte {
	st := (int)(tx.Index)
	ed := (int)(tx.Index) + (int)(tx.Width * tx.Height) - 1

	return basebuf[st:ed]
}

func (tx *Texture) GetPalBuf(basebuf []byte) []byte {
	st := (int)(tx.Index) + (int)(tx.Width * tx.Height)
	ed := (int)(tx.Index) + (int)(tx.Width * tx.Height) + (256 * 3)

	return basebuf[st:ed]
}

func (tx *Texture) GetRgb( pixels []byte, pals []byte, pxidx int) *Rgb {
	r := (*Rgb)(unsafe.Pointer(&pals[int(pixels[pxidx]) * 3]))

	return r
}

func (tx *Texture) UploadTexture(basebuf []byte) {

	var outwidth int = 1
	for outwidth < int(tx.Width) {
		outwidth <<= 1
	}

	if outwidth > 256 {
		outwidth = 256
	}

	var outheight int = 1
	for outheight < int(tx.Height) {
		outheight <<= 1
	}

	if outheight > 256 {
		outheight = 256
	}

	var	row1, row2, col1, col2 [256]int

	for i := 0; i < outwidth; i++ {
		col1[i] = int((float32(i) + 0.25) * (float32(tx.Width) / float32(outwidth)))
		col2[i] = int((float32(i) + 0.75) * (float32(tx.Width) / float32(outwidth)))
	}

	for i := 0; i < outheight; i++ {
		row1[i] = (int)((float32(i) + 0.25) * (float32(tx.Height) / float32(outheight))) * int(tx.Width)
		row2[i] = (int)((float32(i) + 0.75) * (float32(tx.Height) / float32(outheight))) * int(tx.Width)
	}

	pixels := tx.GetPixelBuf(basebuf)
	pals := tx.GetPalBuf(basebuf)
	out := make([]byte, outwidth * outheight * 4)

	// scale down and convert to 32bit RGB
	var out_idx int = 0
	for i := 0; i < outheight; i++ {
		for j := 0; j < outwidth; j++ {
			pix1 := tx.GetRgb(pixels, pals, row1[i] + col1[j])
			pix2 := tx.GetRgb(pixels, pals, row1[i] + col2[j])
			pix3 := tx.GetRgb(pixels, pals, row2[i] + col1[j])
			pix4 := tx.GetRgb(pixels, pals, row2[i] + col2[j])

			out[out_idx + 0] = (pix1.r + pix2.r + pix3.r + pix4.r) >> 2
			out[out_idx + 1] = (pix1.g + pix2.g + pix3.g + pix4.g) >> 2
			out[out_idx + 2] = (pix1.b + pix2.b + pix3.b + pix4.b) >> 2
			out[out_idx + 3] = 0xFF;
			out_idx += 4
		}
	}

	gl.BindTexture(gl.TEXTURE_2D, uint32(g_texnum))
	gl.TexImage2D(gl.TEXTURE_2D, 0, 3, int32(outwidth), int32(outheight), 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(&out[0]))
	gl.TexEnvf(gl.TEXTURE_ENV, gl.TEXTURE_ENV_MODE, gl.MODULATE)
	gl.TexParameterf(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameterf(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)

	tx.Index = g_texnum

	g_texnum++
}
