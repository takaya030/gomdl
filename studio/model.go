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

func (md *Model) GetMesh(basebuf *byte, idx int) *Mesh {
	pmh := (*Mesh)(unsafe.Add(unsafe.Pointer(basebuf), (int)(md.MeshIndex) + (int)(unsafe.Sizeof(Mesh{})) * idx))

	return pmh
}

func (md *Model) GetVertBone(basebuf *byte, idx int) byte {
	pvb := (*byte)(unsafe.Add(unsafe.Pointer(basebuf), (int)(md.VertInfoIndex) + idx))

	return *pvb
}

func (md *Model) GetNormBone(basebuf *byte, idx int) byte {
	pnb := (*byte)(unsafe.Add(unsafe.Pointer(basebuf), (int)(md.NormInfoIndex) + idx))

	return *pnb
}

func (md *Model) GetStudioVert(basebuf *byte, idx int) *Vec3 {
	psv := (*Vec3)(unsafe.Add(unsafe.Pointer(basebuf), (int)(md.VertIndex) + (int)(unsafe.Sizeof(Vec3{})) * idx))

	return psv
}

func (md *Model) GetStudioNorm(basebuf *byte, idx int) *Vec3 {
	psn := (*Vec3)(unsafe.Add(unsafe.Pointer(basebuf), (int)(md.NormIndex) + (int)(unsafe.Sizeof(Vec3{})) * idx))

	return psn
}
