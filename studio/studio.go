package studio

import (
	//"github.com/go-gl/mathgl/mgl32"
)

// lighting options (set Texture.Flags)
const (
	STUDIO_NF_FLATSHADE  = 0x0001
	STUDIO_NF_CHROME     = 0x0002
	STUDIO_NF_FULLBRIGHT = 0x0004
	STUDIO_NF_NOMIPS     = 0x0008
	STUDIO_NF_ALPHA      = 0x0010
	STUDIO_NF_ADDITIVE   = 0x0020
	STUDIO_NF_MASKED     = 0x0040
)

// motion flags (set SeqDesc.MotionType)
const (
	STUDIO_X     = 0x0001
	STUDIO_Y     = 0x0002
	STUDIO_Z     = 0x0004
	STUDIO_XR    = 0x0008
	STUDIO_YR    = 0x0010
	STUDIO_ZR    = 0x0020
	STUDIO_LX    = 0x0040
	STUDIO_LY    = 0x0080
	STUDIO_LZ    = 0x0100
	STUDIO_AX    = 0x0200
	STUDIO_AY    = 0x0400
	STUDIO_AZ    = 0x0800
	STUDIO_AXR   = 0x1000
	STUDIO_AYR   = 0x2000
	STUDIO_AZR   = 0x4000
	STUDIO_TYPES = 0x7FFF
	STUDIO_RLOOP = 0x8000 // controller that wraps shortest distance
)

// size of struct
// (Cannot use unsafe package on GAE)
const (
	SIZEOF_ANIM      = 12
	SIZEOF_ANIMVALUE = 2
)

type Vec3 [3]float32
type Vec4 [4]float32
type Mat34	[3][4]float32

type Bone struct {
	Name           [32]byte   // bone name for symbolic links
	Parent         int32      // parent bone
	Flags          int32      // ??
	BoneController [6]int32   // bone controller index, -1 == none
	Value          [6]float32 // default DoF values
	Scale          [6]float32 // scale for delta DoF values
}

type BoneController struct {
	Bone  int32 // -1 == 0
	Type  int32 // X, Y, Z, XR, YR, ZR, M
	Start float32
	End   float32
	Rest  int32 // byte index value at rest
	Index int32 // 0-3 user set controller, 4 mouth
}

// intersection boxes
type BBox struct {
	Bone  int32
	Group int32  // intersection group
	BbMin Vec3   // bounding box
	BbMax Vec3
}

// attachment
type Attachment struct {
	Name    [32]byte
	Type    int32
	Bone    int32
	Org     Vec3  // attachment point
	Vectors [3]Vec3
}

// demand loaded sequence groups
type SeqGroup struct {
	Label   [32]byte // textual name
	Name    [64]byte // file name
	Unused1 int32    // was "cache"  - index pointer
	Unused2 int32    // was "data" -  hack for group 0
}
