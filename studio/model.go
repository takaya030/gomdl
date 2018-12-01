package studio

import (
	"unsafe"
)

// studio models
type Model struct {
	Name [64]byte

	Type int32

	BoundingRadius float32

	NumMesh   int32
	MeshIndex int32

	NumVerts      int32 // number of unique vertices
	VertInfoIndex int32 // vertex bone info
	VertIndex     int32 // vertex vec3_t
	NumNorms      int32 // number of unique surface normals
	NormInfoIndex int32 // normal bone info
	NormIndex     int32 // normal vec3_t

	NumGroups  int32 // deformation groups
	GroupIndex int32
}

func (m *Model) GetMeshBuf(buf []byte) []byte {
	s := int(m.MeshIndex)
	e := s + int(unsafe.Sizeof(Mesh{}))*int(m.NumMesh)

	return buf[s:e]
}
