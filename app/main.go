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
		mdls := studio.NewModels(bp.GetModelsBuf(buf), int(bp.NumModels))
		fmt.Printf("%# v\n", pretty.Formatter(mdls))

		// read meshes
		for _, mdl := range mdls {
			mshs := studio.NewMeshes(mdl.GetMeshesBuf(buf), int(mdl.NumMesh))
			fmt.Printf("%# v\n", pretty.Formatter(mshs))

			// read tris
			for _, msh := range mshs {
				tris := studio.NewTris(msh.GetTrisBuf(buf))
				fmt.Printf("%# v\n", pretty.Formatter(tris))
			}
		}
	}
}
