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
}
