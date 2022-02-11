package main

import (
	"fmt"
	"github.com/kr/pretty"
	"io/ioutil"
	//"unsafe"

	"github.com/takaya030/gomdl/mdl"
	"github.com/takaya030/gomdl/studio"
)

func main() {
	buf, rferr := ioutil.ReadFile(`./asset/gsg9.mdl`)
	if rferr != nil {
		fmt.Print(rferr)
		return
	}

	// cast hdr
	/*
	p_Hdr := (*studio.Hdr)(unsafe.Pointer(&buf[0]))
	fmt.Printf("%# v\n", pretty.Formatter(*p_Hdr))
	*/

	// read mdldata
	mdd := mdl.NewMdlData(buf)
	mdm := mdl.NewMdlModel(mdd)
	mdm.InitView()
	mdm.SetBlending(0, 0.0)
	mdm.SetBlending(1, 0.0)
	mdm.AdvanceFrame(0.01)
	mdm.SetupModel(0)
	//fmt.Printf("%# v\n", pretty.Formatter(mdd.GetSeqDesc(0)))
	fmt.Printf("%# v\n", pretty.Formatter(*mdm))

	pos1 := studio.Vec3{ 1.0, 0.0, 0.0 }
	ang1 := studio.Vec3{ 0.0, 0.5, 0.0 }
	var qt1 studio.Vec4
	var mat1 studio.Mat34
	var out1 studio.Vec3

	ang1.AngleQuaternion(&qt1)
	qt1.QuaternionMatrix(&mat1)
	pos1.VectorTransform(&mat1, &out1)

	fmt.Printf("%# v\n", pretty.Formatter(pos1))
	fmt.Printf("%# v\n", pretty.Formatter(out1))
	fmt.Printf("%# v\n", pretty.Formatter(mat1))

	// read seqdescs
	/*
		seq := studio.NewSeqDescs(h.GetSeqsBuf(buf), int(h.NumSeq))
		fmt.Printf("%# v\n", pretty.Formatter(seq))
	*/
}
