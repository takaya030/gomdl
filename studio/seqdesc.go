package studio

import (
	"unsafe"
)

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
	LinearMovement     Vec3
	AutoMovePosIndex   int32
	AutoMoveAngleIndex int32

	BbMin Vec3 // per sequence bounding box
	BbMax Vec3

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

// events
type Event struct {
	Frame   int32
	Event   int32
	Type    int32
	Options [64]byte
}

// animations
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

func (sd *SeqDesc) GetAnim(basebuf *byte, idx int) *Anim {
	pan := (*Anim)(unsafe.Add(unsafe.Pointer(basebuf), int(sd.AnimIndex)+int(unsafe.Sizeof(Anim{}))*idx))

	return pan
}

func (anm *Anim) GetNextAnim(idx int) *Anim {
	pan := (*Anim)(unsafe.Add(unsafe.Pointer(anm), int(unsafe.Sizeof(Anim{}))*idx))

	return pan
}

func (anm *Anim) GetAnimValue(idx int) *AnimValue {
	panv := (*AnimValue)(unsafe.Add(unsafe.Pointer(anm), int(anm.Offset[idx])))

	return panv
}

func (anm *Anim) GetAnimValue2(idx int) *AnimValue2 {
	panv := (*AnimValue2)(unsafe.Add(unsafe.Pointer(anm), int(anm.Offset[idx])))

	return panv
}

func (anv *AnimValue) GetAddedPointer(idx int) *AnimValue {
	panv := (*AnimValue)(unsafe.Add(unsafe.Pointer(anv), int(unsafe.Sizeof(AnimValue{}))*idx))

	return panv
}

func (anv *AnimValue) GetAnimValue2Pointer() *AnimValue2 {
	panv := (*AnimValue2)(unsafe.Pointer(anv))

	return panv
}
