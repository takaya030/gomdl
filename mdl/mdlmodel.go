package mdl

import (
	//"bytes"
	//"encoding/binary"
	//"fmt"
	//"github.com/go-gl/mathgl/mgl32"

	"github.com/chewxy/math32"
	"github.com/takaya030/gomdl/studio"
)

// calc utility
type MdlModel struct {
	mdd	*MdlData

	sequence	int32
	frame 		float32
	controller	[4]uint8
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

func (mm *MdlModel) SetController(icntl int32, flval float32) float32 {
	var bc *studio.BoneController = nil

	// find first controller that matches the index
	for i := 0; i < (int)(mm.mdd.GetNumBoneControllers()); i++ {
		tmpbc := mm.mdd.GetBoneController(i)
		if tmpbc.Index == icntl {
			bc = tmpbc
			break
		}
	}

	if bc == nil {
		return flval
	}

	// wrap 0..360 if it's a rotational controller
	if (bc.Type & (studio.STUDIO_XR | studio.STUDIO_YR | studio.STUDIO_ZR)) != 0 {

		// invert value if end < start
		if bc.End < bc.Start {
			flval = -flval
		}

		// does the controller not wrap?
		if bc.Start + 359.0 >= bc.End {
			if flval > ((bc.Start + bc.End) / 2.0) + 180 {
				flval = flval - 360
			}
			if flval < ((bc.Start + bc.End) / 2.0) - 180 {
				flval = flval + 360
			}
		} else {
			if flval > 360 {
				flval = flval - math32.Floor(flval / 360.0) * 360.0
			} else if flval < 0 {
				flval = flval + math32.Floor((flval / -360.0) + 1) * 360.0
			}
		}
	}

	var setting int = (int)(math32.Floor(255 * (flval - bc.Start) / (bc.End - bc.Start)))

	if setting < 0 {
		setting = 0
	}
	if setting > 255 {
		setting = 255
	}
	mm.controller[icntl] = (uint8)(setting)

	return (float32)(setting) * (1.0 / 255.0) * (bc.End - bc.Start) + bc.Start
}
