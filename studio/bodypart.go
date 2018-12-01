package studio

import (
	"unsafe"
)

// body part index
type BodyPart struct {
	Name       [64]byte
	NumModels  int32
	Base       int32
	ModelIndex int32 // index into models array
}

func (b *BodyPart) GetSubModelsBuf(buf []byte) []byte {
	s := int(b.ModelIndex)
	e := s + int(unsafe.Sizeof(Model{}))*int(b.NumModels)

	return buf[s:e]
}
