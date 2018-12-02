package studio

import (
	"bytes"
	"encoding/binary"
	"fmt"
	//"unsafe"
)

// meshes
type Mesh struct {
	NumTris   int32
	TriIndex  int32
	SkinRef   int32
	NumNorms  int32 // per mesh normals (no use)
	NormIndex int32 // normal vec3_t (no use)
}

func NewMeshes(buf []byte, num int) []Mesh {
	m := make([]Mesh, num)
	r := bytes.NewReader(buf)

	// read meshes
	if err := binary.Read(r, binary.LittleEndian, m); err != nil {
		fmt.Print(err)
		return []Mesh{}
	}

	return m
}
