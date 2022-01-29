package mdl

import (
	//"bytes"
	//"encoding/binary"
	//"fmt"

	"github.com/takaya030/gomdl/studio"
)

// unpacked mdl data
type MdlData struct {
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
