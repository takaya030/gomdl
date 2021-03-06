package studio

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"unsafe"

	"github.com/go-gl/mathgl/mgl32"
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

func NewModels(buf []byte, num int) []Model {
	m := make([]Model, num)
	r := bytes.NewReader(buf)

	// read models
	if err := binary.Read(r, binary.LittleEndian, m); err != nil {
		fmt.Print(err)
		return []Model{}
	}

	return m
}

func (m *Model) GetMeshesBuf(buf []byte) []byte {
	s := int(m.MeshIndex)
	e := s + int(unsafe.Sizeof(Mesh{}))*int(m.NumMesh)

	return buf[s:e]
}

func (m *Model) GetVertInfos(buf []byte) []byte {
	s := int(m.VertInfoIndex)
	e := s + int(m.NumVerts)

	return buf[s:e]
}

func (m *Model) GetVertsBuf(buf []byte) []byte {
	s := int(m.VertIndex)
	e := s + int(unsafe.Sizeof(mgl32.Vec3{}))*int(m.NumVerts)

	return buf[s:e]
}

func (m *Model) GetNormInfos(buf []byte) []byte {
	s := int(m.NormInfoIndex)
	e := s + int(m.NumNorms)

	return buf[s:e]
}

func (m *Model) GetNormsBuf(buf []byte) []byte {
	s := int(m.NormIndex)
	e := s + int(unsafe.Sizeof(mgl32.Vec3{}))*int(m.NumNorms)

	return buf[s:e]
}
