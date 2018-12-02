package main

import (
	"fmt"
	"github.com/kr/pretty"
	"io/ioutil"

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

	// read models
	for _, bp := range bps {
		m := studio.NewModels(bp.GetModelsBuf(buf), int(bp.NumModels))
		fmt.Printf("%# v\n", pretty.Formatter(m))
	}
}
