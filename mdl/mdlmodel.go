package mdl

import (
	//"bytes"
	//"encoding/binary"
	//"fmt"
	//"github.com/go-gl/mathgl/mgl32"

	//"github.com/takaya030/gomdl/studio"
)

// calc utility
type MdlModel struct {
	mdd	*MdlData

	sequence	int32
	frame 		float32
}

func NewMdlModel(mdd *MdlData) *MdlModel {
	mdm := new(MdlModel)
	mdm.mdd = mdd

	return mdm
}

func (mm *MdlModel) GetSequence() int32 {
	return mm.sequence
}

func (mm *MdlModel) SetSequence(iseq int32) int32 {
	if iseq >= mm.mdd.GetNumSeq() {
		iseq = 0
	} else if iseq < 0 {
		iseq = mm.mdd.GetNumSeq() - 1
	}

	mm.sequence = iseq
	mm.frame = 0

	return mm.sequence
}
