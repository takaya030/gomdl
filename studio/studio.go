package studio

import (
	"github.com/go-gl/mathgl/mgl32"
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
	Group int32      // intersection group
	BbMin mgl32.Vec3 // bounding box
	BbMax mgl32.Vec3
}

// attachment
type Attachment struct {
	Name    [32]byte
	Type    int32
	Bone    int32
	Org     mgl32.Vec3 // attachment point
	Vectors [3]mgl32.Vec3
}

// body part index
type BodyPart struct {
	Name       [64]byte
	NumModels  int32
	Base       int32
	ModelIndex int32 // index into models array
}

// skin info
type Texture struct {
	Name   [64]byte
	Flags  int32
	Width  int32
	Height int32
	Index  int32
}

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

// meshes
type Mesh struct {
	NumTris   int32
	TriIndex  int32
	SkinRef   int32
	NumNorms  int32 // per mesh normals
	NormIndex int32 // normal vec3_t
}

// sequence descriptions
type SeqDesc struct {
	Label [32]byte // sequence label

	Fps   float32 // frames per second
	Flags int32   // looping/non-looping flags

	Activity  int32
	ActWeight int32

	NumEvents  int32
	EventIndex int32

	NumFrames int32 // number of frames per sequence

	NumPivots  int32 // number of foot pivots
	PivotIndex int32

	MotionType         int32
	MotionBone         int32
	LinearMovement     mgl32.Vec3
	AutoMovePosIndex   int32
	AutoMoveAngleIndex int32

	BbMin mgl32.Vec3 // per sequence bounding box
	BbMax mgl32.Vec3

	NumBlends int32
	AnimIndex int32 // mstudioanim_t pointer relative to start of sequence group data [blend][bone][X, Y, Z, XR, YR, ZR]

	BlendType   [2]int32   // X, Y, Z, XR, YR, ZR
	BlendStart  [2]float32 // starting value
	BlendEnd    [2]float32 // ending value
	BlendParent int32

	SeqGroup int32 // sequence group for demand loading

	EntryNode int32 // transition node at entry
	ExitNode  int32 // transition node at exit
	NodeFlags int32 // transition rules

	NextSeq int32 // auto advancing sequences
}

// demand loaded sequence groups
type SeqGroup struct {
	Label   [32]byte // textual name
	Name    [64]byte // file name
	Unused1 int32    // was "cache"  - index pointer
	Unused2 int32    // was "data" -  hack for group 0
}

type Anim struct {
	Offset [6]uint16
}

// animation frames
type AnimValue struct {
	Valid byte
	Total byte
}

type AnimValue2 struct {
	Value int16
}
