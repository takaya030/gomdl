package studio

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"unsafe"

	//"github.com/go-gl/mathgl/mgl32"
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

func NewSeqDescs(buf []byte, num int) []SeqDesc {
	s := make([]SeqDesc, num)
	r := bytes.NewReader(buf)

	// read seqdescs
	if err := binary.Read(r, binary.LittleEndian, s); err != nil {
		fmt.Print(err)
		return []SeqDesc{}
	}

	return s
}

func (seq *SeqDesc) GetEventsBuf(buf []byte) []byte {
	s := int(seq.EventIndex)
	e := s + int(unsafe.Sizeof(Event{}))*int(seq.NumEvents)

	return buf[s:e]
}

func (seq *SeqDesc) GetAnimBuf(buf []byte, numbones int) []byte {
	s := int(seq.AnimIndex)
	e := s + int(unsafe.Sizeof(Anim{}))*int(seq.NumBlends)*numbones

	return buf[s:e]
}
