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

	ang1 := studio.Vec3{ 0.0, 0.0, 0.0 }
	ang2 := studio.Vec3{ 0.5, 0.5, 0.5 }
	var qt1,qt2,qt3 studio.Vec4
	ang1.AngleQuaternion(&qt1)
	ang2.AngleQuaternion(&qt2)
	qt1.QuaternionSlerp(qt2, 0.1, &qt3)
	fmt.Printf("%# v\n", pretty.Formatter(qt1))
	fmt.Printf("%# v\n", pretty.Formatter(qt2))
	fmt.Printf("%# v\n", pretty.Formatter(qt3))

	// read seqdescs
	/*
		seq := studio.NewSeqDescs(h.GetSeqsBuf(buf), int(h.NumSeq))
		fmt.Printf("%# v\n", pretty.Formatter(seq))
	*/
}
