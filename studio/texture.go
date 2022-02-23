package studio

import (
	//"github.com/go-gl/mathgl/mgl32"
)

// skin info
type Texture struct {
	Name   [64]byte
	Flags  int32
	Width  int32
	Height int32
	Index  int32
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
