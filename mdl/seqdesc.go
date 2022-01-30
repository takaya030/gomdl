package mdl

import (
	//"bytes"
	//"encoding/binary"
	//"fmt"

	"github.com/takaya030/gomdl/studio"
)

// unpacked seqdesc
type SeqDesc struct {
	Seq studio.SeqDesc

	Anims [][]studio.Anim
}

/*
func NewSeqDesc(buf []byte, sd *studio.SeqDesc, numbones int) *SeqDesc {
	s := new(SeqDesc)

	s.Seq = *sd

	// read anims
	s.Anims = make([][]studio.Anim, int(sd.NumBlends))
	r := bytes.NewReader(sd.GetAnimBuf(buf, numbones))

	for i := 0; i < int(sd.NumBlends); i++ {
		s.Anims[i] = make([]studio.Anim, numbones)

		if err := binary.Read(r, binary.LittleEndian, s.Anims[i]); err != nil {
			fmt.Println(err)
			return s
		}
	}

	return s
}
*/
