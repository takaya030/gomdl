package studio

import (
	//"bytes"
	//"encoding/binary"
	//"fmt"
	"unsafe"

	//"github.com/go-gl/mathgl/mgl32"
)

type Hdr struct {
	Id      int32
	Version int32
	Name    [64]byte
	Length  int32

	EyePosition Vec3 // ideal eye position
	Min         Vec3 // ideal movement hull size
	Max         Vec3

	BbMin Vec3 // clipping bounding box
	BbMax Vec3

	Flags int32

	NumBones  int32 // bones
	BoneIndex int32

	NumBoneControllers  int32 // bone controllers
	BoneControllerIndex int32

	NumHitBoxes int32 // complex bounding boxes
	HitBoxIndex int32

	NumSeq   int32 // animation sequences
	SeqIndex int32

	NumSeqGroups  int32 // demand loaded sequences
	SeqGroupIndex int32

	NumTextures      int32 // raw textures
	TextureIndex     int32
	TextureDataIndex int32

	NumSkinRef      int32 // replaceable textures
	NumSkinFamilies int32
	SkinIndex       int32

	NumBodyParts  int32
	BodyPartIndex int32

	NumAttachments  int32 // queryable attachable points
	AttachmentIndex int32

	SoundTable      int32
	SoundIndex      int32
	SoundGroups     int32
	SoundGroupIndex int32

	NumTransitions  int32 // animation node to animation node transition graph
	TransitionIndex int32
}

func NewHdr(buf []byte) *Hdr {
	h := (*Hdr)(unsafe.Pointer(&buf[0]))

	return h
}

func (h *Hdr) GetBonesPtr(buf []byte) *byte {
	pb := (*byte)(unsafe.Add(unsafe.Pointer(&buf[0]), h.BoneIndex))

	return pb
}

func (h *Hdr) GetBoneControllersPtr(buf []byte) *byte {
	pb := (*byte)(unsafe.Add(unsafe.Pointer(&buf[0]), h.BoneControllerIndex))

	return pb
}

func (h *Hdr) GetHitBoxesPtr(buf []byte) *byte {
	pb := (*byte)(unsafe.Add(unsafe.Pointer(&buf[0]), h.HitBoxIndex))

	return pb
}

func (h *Hdr) GetSeqsPtr(buf []byte) *byte {
	pb := (*byte)(unsafe.Add(unsafe.Pointer(&buf[0]), h.SeqIndex))

	return pb
}

func (h *Hdr) GetSeqGroupsPtr(buf []byte) *byte {
	pb := (*byte)(unsafe.Add(unsafe.Pointer(&buf[0]), h.SeqGroupIndex))

	return pb
}

func (h *Hdr) GetTexturesPtr(buf []byte) *byte {
	pb := (*byte)(unsafe.Add(unsafe.Pointer(&buf[0]), h.TextureIndex))

	return pb
}

func (h *Hdr) GetTextureDataPtr(buf []byte) *byte {
	pb := (*byte)(unsafe.Add(unsafe.Pointer(&buf[0]), h.TextureDataIndex))

	return pb
}

func (h *Hdr) GetSkinRefPtr(buf []byte) *byte {
	pb := (*byte)(unsafe.Add(unsafe.Pointer(&buf[0]), h.SkinIndex))

	return pb
}

func (h *Hdr) GetBodyPartsPtr(buf []byte) *byte {
	pb := (*byte)(unsafe.Add(unsafe.Pointer(&buf[0]), h.BodyPartIndex))

	return pb
}

func (h *Hdr) GetAttachmentsPtr(buf []byte) *byte {
	pb := (*byte)(unsafe.Add(unsafe.Pointer(&buf[0]), h.AttachmentIndex))

	return pb
}
