package studio

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"unsafe"
)

// body part index
type BodyPart struct {
	Name       [64]byte
	NumModels  int32
	Base       int32
	ModelIndex int32 // index into models array
}

func NewBodyParts(buf []byte, num int) []BodyPart {
	b := make([]BodyPart, num)
	r := bytes.NewReader(buf)

	// read bodyparts
	if err := binary.Read(r, binary.LittleEndian, b); err != nil {
		fmt.Print(err)
		return []BodyPart{}
	}

	return b
}

func (b *BodyPart) GetModelsBuf(buf []byte) []byte {
	s := int(b.ModelIndex)
	e := s + int(unsafe.Sizeof(Model{}))*int(b.NumModels)

	return buf[s:e]
}
