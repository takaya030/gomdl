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

func (b *BodyPart) GetModel(basebuf *byte, idx int) *Model {
	pmd := (*Model)(unsafe.Add(unsafe.Pointer(basebuf), (int)(b.ModelIndex) + (int)(unsafe.Sizeof(Model{})) * idx))

	return pmd
}
