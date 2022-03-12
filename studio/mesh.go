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
	ptc := (*byte)(unsafe.Add(unsafe.Pointer(basebuf), (int)(me.TriIndex) + int(unsafe.Sizeof(a)) * idx)

	return ptc
}
