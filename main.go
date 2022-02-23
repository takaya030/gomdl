package main

import (
	"fmt"
	"github.com/kr/pretty"
	"io/ioutil"

	"github.com/takaya030/gomdl/mdl"
	"github.com/takaya030/gomdl/studio"
)

func main() {
	buf, rferr := ioutil.ReadFile(`./asset/gsg9.mdl`)
	if rferr != nil {
		fmt.Print(rferr)
		return
	}

	// read mdldata
	mdd := mdl.NewMdlData(buf)
	mdm := mdl.NewMdlModel(mdd)
	mdm.InitView()
	mdm.SetBlending(0, 0.0)
	mdm.SetBlending(1, 0.0)
	mdm.AdvanceFrame(0.01)
	mdm.SetupModel(0)
	fmt.Printf("%# v\n", pretty.Formatter(*(mdd.Hdr)))

	var tex [3]*studio.Texture
	tex[0] = mdd.GetTexture(0)
	tex[1] = mdd.GetTexture(1)
	tex[2] = mdd.GetTexture(2)
	fmt.Printf("%# v\n", pretty.Formatter(tex))

	vec1 := studio.Vec3{ 1.0, 1.0, 0.0 }
	vec2 := studio.Vec3{ 2.0, 0.0, 0.0 }

	vec1.VectorNormalize()
	vec2.VectorNormalize()

	fmt.Printf("%# v\n", pretty.Formatter(vec1))
	fmt.Printf("%# v\n", pretty.Formatter(vec2))
}
