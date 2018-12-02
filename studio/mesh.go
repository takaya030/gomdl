package studio

import (
	"unsafe"

	"github.com/go-gl/mathgl/mgl32"
)

// meshes
type Mesh struct {
	NumTris   int32
	TriIndex  int32
	SkinRef   int32
	NumNorms  int32 // per mesh normals
	NormIndex int32 // normal vec3_t
}

func (m *Mesh) GetNormsBuf(buf []byte) []byte {
	s := int(m.NormIndex)
	e := s + int(unsafe.Sizeof(mgl32.Vec3{}))*int(m.NumNorms)

	return buf[s:e]
}
