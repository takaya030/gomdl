package main

import (
	"fmt"
	"github.com/kr/pretty"
	"io/ioutil"

	"github.com/takaya030/gomdl/mdl"
)

func main() {
	buf, rferr := ioutil.ReadFile(`../asset/sas.mdl`)
	if rferr != nil {
		fmt.Print(rferr)
		return
	}

	// read mdldata
	md := mdl.NewMdlData(buf)
	fmt.Printf("%# v\n", pretty.Formatter(*md))

	// read seqdescs
	/*
		seq := studio.NewSeqDescs(h.GetSeqsBuf(buf), int(h.NumSeq))
		fmt.Printf("%# v\n", pretty.Formatter(seq))
	*/
}
