package mdl

import (
	"github.com/takaya030/gomdl/studio"
)

// unpacked mesh
type Mesh struct {
	SkinRef int32
	Tris    []studio.Tri
}

/*
func NewMesh(buf []byte, msh *studio.Mesh) *Mesh {
	m := new(Mesh)

	m.SkinRef = msh.SkinRef

	// read tris
	m.Tris = studio.NewTris(msh.GetTrisBuf(buf))

	return m
}
*/