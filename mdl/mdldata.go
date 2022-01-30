package mdl

import (
	//"bytes"
	//"encoding/binary"
	//"fmt"
	"unsafe"

	"github.com/takaya030/gomdl/studio"
)

// unpacked mdl data
type MdlData struct {
	BaseBuf         *byte
	Hdr             *studio.Hdr
	Bones           *byte
	BoneControllers *byte
	BBoxes          *byte
	SeqDescs        *byte
	SeqGroups       *byte
	BodyParts       *byte
	Attachments     *byte
	Textures        *byte
	SkinRefs        *byte
}

func NewMdlData(buf []byte) *MdlData {
	md := new(MdlData)

	// base buffer
	md.BaseBuf = (*byte)(unsafe.Pointer(&buf[0]))

	// read hdr
	h := studio.NewHdr(buf)
	md.Hdr = h

	// read bones
	md.Bones = h.GetBonesPtr(buf)

	// read bonecontrollers
	md.BoneControllers = h.GetBoneControllersPtr(buf)

	// read bboxes
	md.BBoxes = h.GetHitBoxesPtr(buf)

	// read seqdesc
	md.SeqDescs = h.GetSeqsPtr(buf)

	// read seqgroups
	md.SeqGroups = h.GetSeqGroupsPtr(buf)

	// read bodyparts
	md.BodyParts = h.GetBodyPartsPtr(buf)

	// read attachments
	md.Attachments = h.GetAttachmentsPtr(buf)

	// read textures
	md.Textures = h.GetTexturesPtr(buf)

	// read skinrefs
	md.SkinRefs = h.GetSkinRefPtr(buf)

	return md
}

func (md *MdlData) GetBone(idx int) *studio.Bone {
	pbn := (*studio.Bone)(unsafe.Add(unsafe.Pointer(md.Bones), (int)(unsafe.Sizeof(studio.Bone{})) * idx))

	return pbn
}

func (md *MdlData) GetBoneController(idx int) *studio.BoneController {
	pbc := (*studio.BoneController)(unsafe.Add(unsafe.Pointer(md.BoneControllers), (int)(unsafe.Sizeof(studio.BoneController{})) * idx))

	return pbc
}

func (md *MdlData) GetBBox(idx int) *studio.BBox {
	pbb := (*studio.BBox)(unsafe.Add(unsafe.Pointer(md.BBoxes), (int)(unsafe.Sizeof(studio.BBox{})) * idx))

	return pbb
}

func (md *MdlData) GetSeqDesc(idx int) *studio.SeqDesc {
	psq := (*studio.SeqDesc)(unsafe.Add(unsafe.Pointer(md.SeqDescs), (int)(unsafe.Sizeof(studio.SeqDesc{})) * idx))

	return psq
}

func (md *MdlData) GetSeqGroup(idx int) *studio.SeqGroup {
	psg := (*studio.SeqGroup)(unsafe.Add(unsafe.Pointer(md.SeqGroups), (int)(unsafe.Sizeof(studio.SeqGroup{})) * idx))

	return psg
}

func (md *MdlData) GetTexture(idx int) *studio.Texture {
	ptx := (*studio.Texture)(unsafe.Add(unsafe.Pointer(md.Textures), (int)(unsafe.Sizeof(studio.Texture{})) * idx))

	return ptx
}