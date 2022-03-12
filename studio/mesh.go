package studio

import (
	//"bytes"
	//"encoding/binary"
	//"fmt"
	"unsafe"
)

// meshes
type Mesh struct {
	NumTris   int32
	TriIndex  int32
	SkinRef   int32
	NumNorms  int32 // per mesh normals (no use)
	NormIndex int32 // normal vec3_t (no use)
}

func (me *Mesh) GetTricmd(basebuf *byte, idx int) *int16 {
	var a int16
	ptc := (*int16)(unsafe.Add(unsafe.Pointer(basebuf), (int)(me.TriIndex) + int(unsafe.Sizeof(a)) * idx))

	return ptc
}

func (me *Mesh) GetNextTricmd(tc *int16, idx int) *int16 {
	var a int16
	p := (*int16)(unsafe.Add(unsafe.Pointer(tc), int(unsafe.Sizeof(a)) * idx))

	return p
}

func (me *Mesh) GetTricmdArray(tc *int16) *[4]int16 {
	p := (*[4]int16)(unsafe.Pointer(tc))

	return p
}
