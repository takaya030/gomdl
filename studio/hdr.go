package studio

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"unsafe"

	"github.com/go-gl/mathgl/mgl32"
)

type Hdr struct {
	Id      int32
	Version int32
	Name    [64]byte
	Length  int32

	EyePosition mgl32.Vec3 // ideal eye position
	Min         mgl32.Vec3 // ideal movement hull size
	Max         mgl32.Vec3

	BbMin mgl32.Vec3 // clipping bounding box
	BbMax mgl32.Vec3

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
	h := new(Hdr)
	r := bytes.NewReader(buf)

	// read hdr
	if err := binary.Read(r, binary.LittleEndian, h); err != nil {
		fmt.Print(err)
		return nil
	}

	return h
}

func (h *Hdr) GetBonesBuf(buf []byte) []byte {
	s := int(h.BoneIndex)
	e := s + int(unsafe.Sizeof(Bone{}))*int(h.NumBones)

	return buf[s:e]
}

func (h *Hdr) GetBoneControllersBuf(buf []byte) []byte {
	s := int(h.BoneControllerIndex)
	e := s + int(unsafe.Sizeof(BoneController{}))*int(h.NumBoneControllers)

	return buf[s:e]
}

func (h *Hdr) GetHitBoxesBuf(buf []byte) []byte {
	s := int(h.HitBoxIndex)
	e := s + int(unsafe.Sizeof(BBox{}))*int(h.NumHitBoxes)

	return buf[s:e]
}

func (h *Hdr) GetSeqBuf(buf []byte) []byte {
	s := int(h.SeqIndex)
	e := s + int(unsafe.Sizeof(SeqDesc{}))*int(h.NumSeq)

	return buf[s:e]
}

func (h *Hdr) GetSeqGroupsBuf(buf []byte) []byte {
	s := int(h.SeqGroupIndex)
	e := s + int(unsafe.Sizeof(SeqGroup{}))*int(h.NumSeqGroups)

	return buf[s:e]
}

func (h *Hdr) GetTexturesBuf(buf []byte) []byte {
	s := int(h.TextureIndex)
	e := s + int(unsafe.Sizeof(Texture{}))*int(h.NumTextures)

	return buf[s:e]
}

func (h *Hdr) GetSkinRefBuf(buf []byte) []byte {
	s := int(h.SkinIndex)
	e := s + int(unsafe.Sizeof(int16(0)))*int(h.NumSkinRef)

	return buf[s:e]
}

func (h *Hdr) GetBodyPartsBuf(buf []byte) []byte {
	s := int(h.BodyPartIndex)
	e := s + int(unsafe.Sizeof(BodyPart{}))*int(h.NumBodyParts)

	return buf[s:e]
}

func (h *Hdr) GetAttachmentsBuf(buf []byte) []byte {
	s := int(h.AttachmentIndex)
	e := s + int(unsafe.Sizeof(Attachment{}))*int(h.NumAttachments)

	return buf[s:e]
}
