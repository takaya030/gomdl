package mdl

import (
	//"bytes"
	//"encoding/binary"
	//"fmt"
	//"github.com/go-gl/mathgl/mgl32"

	//"github.com/takaya030/gomdl/studio"
)

// unpacked model
/*
type Model struct {
	Name [64]byte

	Type int32

	BoundingRadius float32

	Meshes []Mesh

	VertInfos []byte
	Verts     []mgl32.Vec3
	NormInfos []byte
	Norms     []mgl32.Vec3
}
*/

/*
func NewModel(buf []byte, model *studio.Model) *Model {
	m := new(Model)

	m.Name = model.Name
	m.Type = model.Type
	m.BoundingRadius = model.BoundingRadius

	// read studio.Meshes
	mshs := studio.NewMeshes(model.GetMeshesBuf(buf), int(model.NumMesh))
	// read mdl.Mesh
	for _, msh := range mshs {
		m.Meshes = append(m.Meshes, *NewMesh(buf, &msh))
	}

	m.VertInfos = model.GetVertInfos(buf)

	// read verts
	m.Verts = make([]mgl32.Vec3, model.NumVerts)
	r := bytes.NewReader(model.GetVertsBuf(buf))

	if err := binary.Read(r, binary.LittleEndian, m.Verts); err != nil {
		fmt.Print(err)
		return m
	}

	m.NormInfos = model.GetNormInfos(buf)

	// read norms
	m.Norms = make([]mgl32.Vec3, model.NumNorms)
	r2 := bytes.NewReader(model.GetNormsBuf(buf))

	if err := binary.Read(r2, binary.LittleEndian, m.Norms); err != nil {
		fmt.Print(err)
		return m
	}

	return m
}
*/
