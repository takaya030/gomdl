package mdl

import (
	//"bytes"
	//"encoding/binary"
	//"fmt"
	//"github.com/go-gl/mathgl/mgl32"

	"github.com/chewxy/math32"
	"github.com/takaya030/gomdl/studio"
)

const (
	MAXSTUDIOBONES	= 128		// total bones actually used
)

// global variables

// for lighting
var	g_lightvec		studio.Vec3
var	g_lightcolor	studio.Vec3
var g_blightvec		[MAXSTUDIOBONES]studio.Vec3
var	g_ambientlight	float32		// ambient world light
var	g_shadelight	float32		// direct world light

var	g_bonetransform		[MAXSTUDIOBONES]studio.Mat34	// bone transformation matrix


// calc utility
type MdlModel struct {
	mdd	*MdlData

	// entity settings
	sequence	int32		// sequence index
	frame 		float32		// frame
	controller	[4]uint8	// bone controllers
	mouth		uint8		// mouth position
	blending	[2]uint8	// animation blending
	bodynum		int32		// bodypart selection

	// internal data
	pmodel		*studio.Model
	adj			studio.Vec4
}

func NewMdlModel(mdd *MdlData) *MdlModel {
	mdm := new(MdlModel)
	mdm.mdd = mdd

	return mdm
}

func (mm *MdlModel) InitView() {
	mm.SetSequence(0)
	mm.SetController(0, 0.0)
	mm.SetController(1, 0.0)
	mm.SetController(2, 0.0)
	mm.SetController(3, 0.0)
	mm.SetMouth(0.0)
	mm.bodynum = 0
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

func (mm *MdlModel) SetMouth(flval float32) float32 {
	var bc *studio.BoneController = nil

	// find first controller that matches the index
	for i := 0; i < (int)(mm.mdd.GetNumBoneControllers()); i++ {
		tmpbc := mm.mdd.GetBoneController(i)
		if tmpbc.Index == 4 {
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

	var setting int = (int)(math32.Floor(64 * (flval - bc.Start) / (bc.End - bc.Start)))

	if setting < 0 {
		setting = 0
	}
	if setting > 64 {
		setting = 64
	}
	mm.mouth = (uint8)(setting)

	return (float32)(setting) * (1.0 / 64.0) * (bc.End - bc.Start) + bc.Start
}

func (mm *MdlModel) SetBlending(iblender int32, flval float32) float32 {
	var seq *studio.SeqDesc = mm.mdd.GetSeqDesc((int)(mm.sequence))

	if seq.BlendType[iblender] == 0 {
		return flval
	}

	if (seq.BlendType[iblender] & (studio.STUDIO_XR | studio.STUDIO_YR | studio.STUDIO_ZR)) != 0 {

		// invert value if end < start
		if seq.BlendEnd[iblender] < seq.BlendStart[iblender] {
			flval = -flval
		}

		// does the controller not wrap?
		if seq.BlendStart[iblender] + 359.0 >= seq.BlendEnd[iblender] {
			if flval > ((seq.BlendStart[iblender] + seq.BlendEnd[iblender]) / 2.0) + 180 {
				flval = flval - 360
			}
			if flval < ((seq.BlendStart[iblender] + seq.BlendEnd[iblender]) / 2.0) - 180 {
				flval = flval + 360
			}
		}
	}

	var setting int = (int)(math32.Floor(255 * (flval - seq.BlendStart[iblender]) / (seq.BlendEnd[iblender] - seq.BlendStart[iblender])))

	if setting < 0 {
		setting = 0
	}
	if setting > 255 {
		setting = 255
	}
	mm.blending[iblender] = (uint8)(setting)

	return (float32)(setting) * (1.0 / 255.0) * (seq.BlendEnd[iblender] - seq.BlendStart[iblender]) + seq.BlendStart[iblender]
}

func (mm *MdlModel) AdvanceFrame(dt float32) {
	var seq *studio.SeqDesc = mm.mdd.GetSeqDesc((int)(mm.sequence))

	if dt > 0.1 {
		dt = 0.1
	}

	mm.frame += dt * seq.Fps

	if seq.NumFrames <= 1 {
		mm.frame = 0
	} else {
		// wrap
		mm.frame -= math32.Floor(mm.frame / (float32)(seq.NumFrames - 1)) * (float32)(seq.NumFrames - 1)
	}
}

func (mm *MdlModel) SetupLighting() {
	g_ambientlight = 32.0
	g_shadelight = 192.0

	g_lightvec = studio.Vec3{ 0, 0, -1.0 }

	g_lightcolor = studio.Vec3{ 1.0, 1.0, 1.0 }

	for i := 0; i < (int)(mm.mdd.GetNumBones()); i++ {
		g_lightvec.VectorIRotate( &(g_bonetransform[i]), &(g_blightvec[i]) )
	}

}

func (mm *MdlModel) SetupModel(bodypart int32) {
	if bodypart >= mm.mdd.GetNumBodyParts() {
		bodypart = 0
	}

	pbp := mm.mdd.GetBodyPart((int)(bodypart))

	var index int32 = (mm.bodynum / pbp.Base) % pbp.NumModels

	mm.pmodel = pbp.GetModel(mm.mdd.BaseBuf, (int)(index))
}

func (mm *MdlModel) CalcBoneAdj() {
	var value float32 = 0.0

	for j := 0; j < (int)(mm.mdd.GetNumBoneControllers()); j++ {
		tmpbc := mm.mdd.GetBoneController(j)
		i := (int)(tmpbc.Index)
		if i <= 3 {
			if (tmpbc.Type & studio.STUDIO_RLOOP) != 0 {
				value = (float32)(mm.controller[i]) * (360.0 / 256.0) + tmpbc.Start
			} else {
				value = (float32)(mm.controller[i]) / 255.0
				if value < 0 {
					value = 0
				}
				if value > 1.0 {
					value = 1.0
				}
				value = (1.0 - value) * tmpbc.Start + value * tmpbc.End
			}

		} else {
			value = (float32)(mm.mouth) / 64.0
			if value > 1.0 {
				value = 1.0
			}
			value = (1.0 - value) * tmpbc.Start + value * tmpbc.End
		}

		switch tmpbc.Type & studio.STUDIO_TYPES {
		case studio.STUDIO_XR, studio.STUDIO_YR, studio.STUDIO_ZR:
			mm.adj[j] = value * (math32.Pi / 180.0)
		case studio.STUDIO_X, studio.STUDIO_Y, studio.STUDIO_Z:
			mm.adj[j] = value
		}
	}
}

func (mm *MdlModel) CalcBoneQuaternion(frame int, s float32, pbone *studio.Bone, panim *studio.Anim, q *studio.Vec4) {
	var angle1, angle2 studio.Vec3

	for j := 0; j < 3; j++ {

		if panim.Offset[j+3] == 0 {
			angle1[j] = pbone.Value[j+3]		// default
			angle2[j] = pbone.Value[j+3]		// default
		} else {
			var panimvalue *studio.AnimValue
			var panimvalue2 *studio.AnimValue2

			panimvalue = panim.GetAnimValue(j + 3)
			k := frame
			for (int)(panimvalue.Total) <= k {
				k -= (int)(panimvalue.Total)
				panimvalue = panimvalue.GetAddedPointer((int)(panimvalue.Valid) + 1)
			}
			// Bah, missing blend!
			if (int)(panimvalue.Valid) > k {
				panimvalue2 = panimvalue.GetAddedPointer(k+1).GetAnimValue2Pointer()
				angle1[j] = (float32)(panimvalue2.Value)

				if (int)(panimvalue.Valid) > k + 1 {
					panimvalue2 = panimvalue.GetAddedPointer(k+2).GetAnimValue2Pointer()
					angle2[j] = (float32)(panimvalue2.Value)
				} else {
					if (int)(panimvalue.Total) > k + 1 {
						angle2[j] = angle1[j]
					} else {
						panimvalue2 = panimvalue.GetAddedPointer((int)(panimvalue.Valid) + 2).GetAnimValue2Pointer()
						angle2[j] = (float32)(panimvalue2.Value)
					}
				}
			} else {
				panimvalue2 = panimvalue.GetAddedPointer((int)(panimvalue.Valid)).GetAnimValue2Pointer()
				angle1[j] = (float32)(panimvalue2.Value)
				if (int)(panimvalue.Total) > k + 1 {
					angle2[j] = angle1[j]
				} else {
					panimvalue2 = panimvalue.GetAddedPointer((int)(panimvalue.Valid) + 2).GetAnimValue2Pointer()
					angle2[j] = (float32)(panimvalue2.Value)
				}
			}
			angle1[j] = pbone.Value[j+3] + angle1[j] * pbone.Scale[j+3]
			angle2[j] = pbone.Value[j+3] + angle2[j] * pbone.Scale[j+3]
		}

		if (int)(pbone.BoneController[j+3]) != -1 {
			angle1[j] += mm.adj[(int)(pbone.BoneController[j+3])]
			angle2[j] += mm.adj[(int)(pbone.BoneController[j+3])]
		}
	}

	var	q1, q2 studio.Vec4
	if angle1.VectorCompare(&angle2) == false {
		angle1.AngleQuaternion(&q1)
		angle2.AngleQuaternion(&q2)
		q1.QuaternionSlerp(q2, s, q)
	} else {
		angle1.AngleQuaternion(q)
	}
}
