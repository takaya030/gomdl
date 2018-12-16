package mdl

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/takaya030/gomdl/studio"
)

// unpacked mdl data
type MdlData struct {
	Hdr             *studio.Hdr
	Bones           []studio.Bone
	BoneControllers []studio.BoneController
	BBoxes          []studio.BBox
	SeqDescs        []SeqDesc
	SeqGroups       []studio.SeqGroup
	BodyParts       []BodyPart
	Attachments     []studio.Attachment
	SkinRefs        []int16
	//Textures        []studio.Texture // skin info
}

func NewMdlData(buf []byte) *MdlData {
	md := new(MdlData)

	// read hdr
	h := studio.NewHdr(buf)
	md.Hdr = h

	// read bones
	md.Bones = make([]studio.Bone, int(h.NumBones))
	r := bytes.NewReader(h.GetBonesBuf(buf))

	if err := binary.Read(r, binary.LittleEndian, md.Bones); err != nil {
		fmt.Print(err)
		return md
	}

	// read bonecontrollers
	md.BoneControllers = make([]studio.BoneController, int(h.NumBoneControllers))
	r = bytes.NewReader(h.GetBoneControllersBuf(buf))

	if err := binary.Read(r, binary.LittleEndian, md.BoneControllers); err != nil {
		fmt.Print(err)
		return md
	}

	// read bboxes
	md.BBoxes = make([]studio.BBox, int(h.NumHitBoxes))
	r = bytes.NewReader(h.GetHitBoxesBuf(buf))

	if err := binary.Read(r, binary.LittleEndian, md.BBoxes); err != nil {
		fmt.Print(err)
		return md
	}

	// read seqdesc
	sds := studio.NewSeqDescs(h.GetSeqsBuf(buf), int(h.NumSeq))
	// read mdl.SeqDesc
	for _, sd := range sds {

		md.SeqDescs = append(md.SeqDescs, *NewSeqDesc(buf, &sd, int(h.NumBones)))
	}

	// read seqgroups
	md.SeqGroups = make([]studio.SeqGroup, int(h.NumSeqGroups))
	r = bytes.NewReader(h.GetSeqGroupsBuf(buf))

	if err := binary.Read(r, binary.LittleEndian, md.SeqGroups); err != nil {
		fmt.Print(err)
		return md
	}

	// read bodyparts
	bps := studio.NewBodyParts(h.GetBodyPartsBuf(buf), int(h.NumBodyParts))
	// read mdl.BodyPart
	for _, bp := range bps {

		md.BodyParts = append(md.BodyParts, *NewBodyPart(buf, &bp))
	}

	// read attachments
	md.Attachments = make([]studio.Attachment, int(h.NumAttachments))
	r = bytes.NewReader(h.GetAttachmentsBuf(buf))

	if err := binary.Read(r, binary.LittleEndian, md.Attachments); err != nil {
		fmt.Print(err)
		return md
	}

	// read skinrefs
	md.SkinRefs = make([]int16, int(h.NumSkinRef))
	r = bytes.NewReader(h.GetSkinRefBuf(buf))

	if err := binary.Read(r, binary.LittleEndian, md.SkinRefs); err != nil {
		fmt.Print(err)
		return md
	}

	return md
}
