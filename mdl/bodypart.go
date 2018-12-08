package mdl

import (
	"github.com/takaya030/gomdl/studio"
)

// unpacked bodypart
type BodyPart struct {
	Name [64]byte

	Base int32

	Models []Model
}

func NewBodyPart(buf []byte, sbp *studio.BodyPart) *BodyPart {
	b := new(BodyPart)

	b.Name = sbp.Name
	b.Base = sbp.Base

	// read studio.Models
	models := studio.NewModels(sbp.GetModelsBuf(buf), int(sbp.NumModels))
	for _, model := range models {

		// read mdl.Model
		b.Models = append(b.Models, *NewModel(buf, &model))
	}

	return b
}
