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
	MAXSTUDIOVERTS	= 2048
)

// global variables

var g_vright		studio.Vec3		// needs to be set to viewer's right in order for chrome to work

var g_xformverts	[MAXSTUDIOVERTS]studio.Vec3		// transformed vertices
var g_lightvalues	[MAXSTUDIOVERTS]studio.Vec3		// light surface normals

// for lighting
var	g_lightvec		studio.Vec3
var	g_lightcolor	studio.Vec3
var g_blightvec		[MAXSTUDIOBONES]studio.Vec3
var	g_ambientlight	float32		// ambient world light
var	g_shadelight	float32		// direct world light

var g_smodels_total	int			// cookie

var	g_bonetransform		[MAXSTUDIOBONES]studio.Mat34	// bone transformation matrix

var g_chrome		[MAXSTUDIOVERTS][2]int		// texture coords for surface normals
var g_chromeage		[MAXSTUDIOBONES]int			// last time chrome vectors were updated
var g_chromeup		[MAXSTUDIOBONES]studio.Vec3	// chrome vector "up" in bone reference frames
var g_chromeright	[MAXSTUDIOBONES]studio.Vec3	// chrome vector "right" in bone reference frames

// for setup bones
var w_pos			[MAXSTUDIOBONES]studio.Vec3
var w_bonematrix	studio.Mat34
var w_q				[MAXSTUDIOBONES]studio.Vec4

var w_pos2			[MAXSTUDIOBONES]studio.Vec3
var w_q2			[MAXSTUDIOBONES]studio.Vec4
var w_pos3			[MAXSTUDIOBONES]studio.Vec3
var w_q3			[MAXSTUDIOBONES]studio.Vec4
var w_pos4			[MAXSTUDIOBONES]studio.Vec3
var w_q4			[MAXSTUDIOBONES]studio.Vec4

// calc utility
type MdlModel struct {
	mdd	*MdlData

	// entity settings
	origin		studio.Vec3
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
	for i := 0; i < int(mm.mdd.GetNumBoneControllers()); i++ {
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

	var setting int = int(math32.Floor(255 * (flval - bc.Start) / (bc.End - bc.Start)))

	if setting < 0 {
		setting = 0
	}
	if setting > 255 {
		setting = 255
	}
	mm.controller[icntl] = uint8(setting)

	return float32(setting) * (1.0 / 255.0) * (bc.End - bc.Start) + bc.Start
}

func (mm *MdlModel) SetMouth(flval float32) float32 {
	var bc *studio.BoneController = nil

	// find first controller that matches the index
	for i := 0; i < int(mm.mdd.GetNumBoneControllers()); i++ {
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

	var setting int = int(math32.Floor(64 * (flval - bc.Start) / (bc.End - bc.Start)))

	if setting < 0 {
		setting = 0
	}
	if setting > 64 {
		setting = 64
	}
	mm.mouth = uint8(setting)

	return float32(setting) * (1.0 / 64.0) * (bc.End - bc.Start) + bc.Start
}

func (mm *MdlModel) SetBlending(iblender int32, flval float32) float32 {
	var seq *studio.SeqDesc = mm.mdd.GetSeqDesc(int(mm.sequence))

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

	var setting int = int(math32.Floor(255 * (flval - seq.BlendStart[iblender]) / (seq.BlendEnd[iblender] - seq.BlendStart[iblender])))

	if setting < 0 {
		setting = 0
	}
	if setting > 255 {
		setting = 255
	}
	mm.blending[iblender] = uint8(setting)

	return float32(setting) * (1.0 / 255.0) * (seq.BlendEnd[iblender] - seq.BlendStart[iblender]) + seq.BlendStart[iblender]
}

func (mm *MdlModel) AdvanceFrame(dt float32) {
	var seq *studio.SeqDesc = mm.mdd.GetSeqDesc(int(mm.sequence))

	if dt > 0.1 {
		dt = 0.1
	}

	mm.frame += dt * seq.Fps

	if seq.NumFrames <= 1 {
		mm.frame = 0
	} else {
		// wrap
		mm.frame -= math32.Floor(mm.frame / float32(seq.NumFrames - 1)) * float32(seq.NumFrames - 1)
	}
}

func (mm *MdlModel) SetupLighting() {
	g_ambientlight = 32.0
	g_shadelight = 192.0

	g_lightvec = studio.Vec3{ 0, 0, -1.0 }

	g_lightcolor = studio.Vec3{ 1.0, 1.0, 1.0 }

	for i := 0; i < int(mm.mdd.GetNumBones()); i++ {
		g_lightvec.VectorIRotate( &(g_bonetransform[i]), &(g_blightvec[i]) )
	}

}

func (mm *MdlModel) SetupModel(bodypart int32) {
	if bodypart >= mm.mdd.GetNumBodyParts() {
		bodypart = 0
	}

	pbp := mm.mdd.GetBodyPart(int(bodypart))

	var index int32 = (mm.bodynum / pbp.Base) % pbp.NumModels

	mm.pmodel = pbp.GetModel(mm.mdd.BaseBuf, int(index))
}

func (mm *MdlModel) CalcBoneAdj() {
	var value float32 = 0.0

	for j := 0; j < int(mm.mdd.GetNumBoneControllers()); j++ {
		tmpbc := mm.mdd.GetBoneController(j)
		i := int(tmpbc.Index)
		if i <= 3 {
			if (tmpbc.Type & studio.STUDIO_RLOOP) != 0 {
				value = float32(mm.controller[i]) * (360.0 / 256.0) + tmpbc.Start
			} else {
				value = float32(mm.controller[i]) / 255.0
				if value < 0 {
					value = 0
				}
				if value > 1.0 {
					value = 1.0
				}
				value = (1.0 - value) * tmpbc.Start + value * tmpbc.End
			}

		} else {
			value = float32(mm.mouth) / 64.0
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
			for int(panimvalue.Total) <= k {
				k -= int(panimvalue.Total)
				panimvalue = panimvalue.GetAddedPointer(int(panimvalue.Valid) + 1)
			}
			// Bah, missing blend!
			if int(panimvalue.Valid) > k {
				panimvalue2 = panimvalue.GetAddedPointer(k+1).GetAnimValue2Pointer()
				angle1[j] = float32(panimvalue2.Value)

				if int(panimvalue.Valid) > k + 1 {
					panimvalue2 = panimvalue.GetAddedPointer(k+2).GetAnimValue2Pointer()
					angle2[j] = float32(panimvalue2.Value)
				} else {
					if int(panimvalue.Total) > k + 1 {
						angle2[j] = angle1[j]
					} else {
						panimvalue2 = panimvalue.GetAddedPointer(int(panimvalue.Valid) + 2).GetAnimValue2Pointer()
						angle2[j] = float32(panimvalue2.Value)
					}
				}
			} else {
				panimvalue2 = panimvalue.GetAddedPointer(int(panimvalue.Valid)).GetAnimValue2Pointer()
				angle1[j] = float32(panimvalue2.Value)
				if int(panimvalue.Total) > k + 1 {
					angle2[j] = angle1[j]
				} else {
					panimvalue2 = panimvalue.GetAddedPointer(int(panimvalue.Valid) + 2).GetAnimValue2Pointer()
					angle2[j] = float32(panimvalue2.Value)
				}
			}
			angle1[j] = pbone.Value[j+3] + angle1[j] * pbone.Scale[j+3]
			angle2[j] = pbone.Value[j+3] + angle2[j] * pbone.Scale[j+3]
		}

		if int(pbone.BoneController[j+3]) != -1 {
			angle1[j] += mm.adj[int(pbone.BoneController[j+3])]
			angle2[j] += mm.adj[int(pbone.BoneController[j+3])]
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

func (mm *MdlModel) CalcBonePosition(frame int, s float32, pbone *studio.Bone, panim *studio.Anim, pos *studio.Vec3) {

	for j := 0; j < 3; j++ {
		pos[j] = pbone.Value[j]; // default;
		if panim.Offset[j] != 0 {
			var panimvalue *studio.AnimValue
			var panimvalue2 *studio.AnimValue2
			var panimvalue2_2 *studio.AnimValue2

			panimvalue = panim.GetAnimValue(j)
			k := frame;
			// find span of values that includes the frame we want
			for int(panimvalue.Total) <= k {
				k -= int(panimvalue.Total)
				panimvalue = panimvalue.GetAddedPointer(int(panimvalue.Valid) + 1)
			}
			// if we're inside the span
			if int(panimvalue.Valid) > k {
				// and there's more data in the span
				if int(panimvalue.Valid) > k + 1 {
					panimvalue2 = panimvalue.GetAddedPointer(k+1).GetAnimValue2Pointer()
					panimvalue2_2 = panimvalue.GetAddedPointer(k+2).GetAnimValue2Pointer()
					pos[j] += (float32(panimvalue2.Value) * (1.0 - s) + s * float32(panimvalue2_2.Value)) * pbone.Scale[j]
				} else {
					panimvalue2 = panimvalue.GetAddedPointer(k+1).GetAnimValue2Pointer()
					pos[j] += float32(panimvalue2.Value) * pbone.Scale[j]
				}
			} else {
				// are we at the end of the repeating values section and there's another section with data?
				if int(panimvalue.Total) <= k + 1 {
					panimvalue2 = panimvalue.GetAddedPointer(int(panimvalue.Valid)).GetAnimValue2Pointer()
					panimvalue2_2 = panimvalue.GetAddedPointer(int(panimvalue.Valid)+2).GetAnimValue2Pointer()
					pos[j] += (float32(panimvalue2.Value) * (1.0 - s) + s * float32(panimvalue2_2.Value)) * pbone.Scale[j]
				} else {
					panimvalue2 = panimvalue.GetAddedPointer(int(panimvalue.Valid)).GetAnimValue2Pointer()
					pos[j] += float32(panimvalue2.Value) * pbone.Scale[j]
				}
			}
		}
		if int(pbone.BoneController[j]) != -1 {
			pos[j] += mm.adj[int(pbone.BoneController[j])]
		}
	}
}

func (mm *MdlModel) CalcRotations (pos *[MAXSTUDIOBONES]studio.Vec3, q *[MAXSTUDIOBONES]studio.Vec4, pseqdesc *studio.SeqDesc, panim *studio.Anim, f float32) {

	var frame int = int(f)
	var s float32 = f - float32(frame)

	// add in programatic controllers
	mm.CalcBoneAdj()

	for i := 0; i < int(mm.mdd.GetNumBones()); i++ {
		pbone := mm.mdd.GetBone(i)
		mm.CalcBoneQuaternion( frame, s, pbone, panim, &q[i] )
		mm.CalcBonePosition( frame, s, pbone, panim, &pos[i] )
		panim = panim.GetNextAnim(1)
	}

	if (pseqdesc.MotionType & studio.STUDIO_X) != 0 {
		pos[int(pseqdesc.MotionBone)][0] = 0.0
	}
	if (pseqdesc.MotionType & studio.STUDIO_Y) != 0 {
		pos[int(pseqdesc.MotionBone)][1] = 0.0
	}
	if (pseqdesc.MotionType & studio.STUDIO_Z) != 0 {
		pos[int(pseqdesc.MotionBone)][2] = 0.0
	}
}

func (mm *MdlModel) SlerpBones(q1 *[MAXSTUDIOBONES]studio.Vec4, pos1 *[MAXSTUDIOBONES]studio.Vec3, q2 *[MAXSTUDIOBONES]studio.Vec4, pos2 *[MAXSTUDIOBONES]studio.Vec3, s float32) {

	if s < 0 {
		s = 0
	} else if s > 1.0 {
		s = 1.0
	}

	var s1 float32 = 1.0 - s
	var q3 studio.Vec4

	for i := 0; i < int(mm.mdd.GetNumBones()); i++ {
		q1[i].QuaternionSlerp( q2[i], s, &q3 )
		q1[i] = q3
		pos1[i][0] = pos1[i][0] * s1 + pos2[i][0] * s
		pos1[i][1] = pos1[i][1] * s1 + pos2[i][1] * s
		pos1[i][2] = pos1[i][2] * s1 + pos2[i][2] * s
	}
}

func (mm *MdlModel) SetUpBones() {

	if mm.sequence >=  mm.mdd.GetNumSeq() {
		mm.sequence = 0
	}

	var pseqdesc *studio.SeqDesc = mm.mdd.GetSeqDesc(int(mm.sequence))
	var panim *studio.Anim = pseqdesc.GetAnim(mm.mdd.BaseBuf, 0)
	mm.CalcRotations(&w_pos, &w_q, pseqdesc, panim, mm.frame)

	if pseqdesc.NumBlends > 1 {
		var s float32

		panim = panim.GetNextAnim(int(mm.mdd.GetNumBones()))
		mm.CalcRotations(&w_pos2, &w_q2, pseqdesc, panim, mm.frame)
		s = float32(mm.blending[0]) / 255.0

		mm.SlerpBones(&w_q, &w_pos, &w_q2, &w_pos2, s)

		if pseqdesc.NumBlends == 4 {
			panim = panim.GetNextAnim(int(mm.mdd.GetNumBones()))
			mm.CalcRotations(&w_pos3, &w_q3, pseqdesc, panim, mm.frame)

			panim = panim.GetNextAnim(int(mm.mdd.GetNumBones()))
			mm.CalcRotations(&w_pos4, &w_q4, pseqdesc, panim, mm.frame)

			s = float32(mm.blending[0]) / 255.0
			mm.SlerpBones(&w_q3, &w_pos3, &w_q4, &w_pos4, s)

			s = float32(mm.blending[1]) / 255.0
			mm.SlerpBones(&w_q, &w_pos, &w_q3, &w_pos3, s)
		}
	}

	for i := 0; i < int(mm.mdd.GetNumBones()); i++ {
		pbone := mm.mdd.GetBone(i)

		w_q[i].QuaternionMatrix(&w_bonematrix)

		w_bonematrix[0][3] = w_pos[i][0];
		w_bonematrix[1][3] = w_pos[i][1];
		w_bonematrix[2][3] = w_pos[i][2];

		if pbone.Parent == -1 {
			g_bonetransform[i] = w_bonematrix
		} else {
			g_bonetransform[int(pbone.Parent)].ConcatTransforms (&w_bonematrix, &g_bonetransform[i])
		}
	}
}

func (mm *MdlModel) Lighting(bone int, flags int32, normal *Vec3) float32 {
	var illum float32
	var lightcos float32

	illum = g_ambientlight

	if (flags & studio.STUDIO_NF_FLATSHADE) != 0 {
		illum += g_shadelight * 0.8
	} else {
		lightcos = normal.DotProduct(&g_blightvec[bone])	// -1 colinear, 1 opposite

		if lightcos > 1.0 {
			lightcos = 1.0
		}

		illum += g_shadelight

		var r float32 = g_lambert
		if r <= 1.0 {
			r = 1.0
		}

		lightcos = (lightcos + (r - 1.0)) / r 		// do modified hemispherical lighting
		if lightcos > 0.0 {
			illum -= g_shadelight * lightcos
		}
		if (illum <= 0) {
			illum = 0
		}
	}

	if illum > 255.0 {
		illum = 255.0
	}

	return illum / 255.0	// Light from 0 to 1.0
}

func (mm *MdlModel) Chrome(pchrome *[2]int, bone int, normal *studio.Vec3) {

	if g_chromeage[bone] != g_smodels_total {
		// calculate vectors from the viewer to the bone. This roughly adjusts for position
		var chromeupvec		studio.Vec3	// g_chrome t vector in world reference frame
		var chromerightvec	studio.Vec3	// g_chrome s vector in world reference frame
		var tmp				studio.Vec3	// vector pointing at bone in world reference frame

		mm.origin.VectorScale(-1.0, &tmp)
		tmp[0] += g_bonetransform[bone][0][3]
		tmp[1] += g_bonetransform[bone][1][3]
		tmp[2] += g_bonetransform[bone][2][3]
		tmp.VectorNormalize()
		tmp.CrossProduct(&g_vright, &chromeupvec)
		chromeupvec.VectorNormalize()
		tmp.CrossProduct(&chromeupvec, &chromerightvec)
		chromerightvec.VectorNormalize()

		chromeupvec.VectorIRotate(&g_bonetransform[bone], &g_chromeup[bone])
		chromerightvec.VectorIRotate(&g_bonetransform[bone], &g_chromeright[bone])

		g_chromeage[bone] = g_smodels_total
	}

	var	n float32

	// calc s coord
	n = normal.DotProduct(&g_chromeright[bone])
	pchrome[0] = (n + 1.0) * 32.0		// FIX: make this a float

	// calc t coord
	n = normal.DotProduct(&g_chromeup[bone])
	pchrome[1] = (n + 1.0) * 32.0		// FIX: make this a float
}

func (mm *MdlModel) DrawPoints () {

	for i := 0; i < int(mm.pmodel.NumVerts); i++ {
		pstudiovert := mm.pmodel.GetStudioVert(mm.mdd.BaseBuf, i)
		vertbone := int(mm.pmodel.GetVertBone(mm.mdd.BaseBuf, i))
		pstudiovert.VectorTransform(&g_bonetransform[vertbone], &g_xformverts[i])
	}

	var lv_idx int = 0
	var norm_idx int = 0
	for j := 0; j < int(mm.pmodel.NumMesh); j++ {
		pmesh := mm.pmodel.GetMesh(mm.mdd.BaseBuf, j)
		pskinref := mm.mdd.GetSkinRef(int(pmesh.SkinRef))
		ptexture := mm.mdd.GetTexture(int(*pskinref))
		flags := ptexture.Flags

		for i := 0; i < int(pmesh.NumNorms); i++ {
			pstudionorm := mm.pmodel.GetStudioNorm(mm.mdd.BaseBuf, norm_idx)
			normbone := int(mm.pmodel.GetNormBone(mm.mdd.BaseBuf, norm_idx))
			lv_tmp := mm.Lighting(normbone, flags, pstudionorm)

			if (flags & STUDIO_NF_CHROME) != 0 {
				mm.Chrome( &g_chrome[lv_idx], normbone, pstudionorm )
			}

			g_lightvalues[lv_idx][0] = lv_tmp * g_lightcolor[0]
			g_lightvalues[lv_idx][1] = lv_tmp * g_lightcolor[1]
			g_lightvalues[lv_idx][2] = lv_tmp * g_lightcolor[2]

			lv_idx++
			norm_idx++
		}

	}
}
