package main

import (
	"fmt"
	"github.com/kr/pretty"
	"io/ioutil"

	"github.com/takaya030/gomdl/mdl"
	"github.com/takaya030/gomdl/studio"
)

func main() {
	buf, rferr := ioutil.ReadFile(`../asset/sas.mdl`)
	if rferr != nil {
		fmt.Print(rferr)
		return
	}

	// read hdr
	h := studio.NewHdr(buf)
	fmt.Printf("%# v\n", pretty.Formatter(*h))

	// read bodyparts
	bps := studio.NewBodyParts(h.GetBodyPartsBuf(buf), int(h.NumBodyParts))
	fmt.Printf("%# v\n", pretty.Formatter(bps))

	// read mdl.BodyPart
	for _, bp := range bps {

		mdlbodypart := mdl.NewBodyPart( buf, &bp )
		fmt.Printf("%# v\n", pretty.Formatter(*mdlbodypart))
	}

	// read seqdescs
	/*
		seq := studio.NewSeqDescs(h.GetSeqsBuf(buf), int(h.NumSeq))
		fmt.Printf("%# v\n", pretty.Formatter(seq))
	*/
}
