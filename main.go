package main

import (
	"fmt"
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	var winTitle string = "Go-SDL2 + Go-GL"
	var winWidth, winHeight int32 = 800, 600
	var window *sdl.Window
	var context sdl.GLContext
	var event sdl.Event
	var running bool
	var err error

	if err = sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err = sdl.CreateWindow(winTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		winWidth, winHeight, sdl.WINDOW_OPENGL)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()
	context, err = window.GLCreateContext()
	if err != nil {
		panic(err)
	}
	defer sdl.GLDeleteContext(context)

	if err = gl.Init(); err != nil {
		panic(err)
	}

	gl.Enable(gl.DEPTH_TEST)
	gl.ClearColor(0.2, 0.2, 0.3, 1.0)
	gl.ClearDepth(1)
	gl.DepthFunc(gl.LEQUAL)
	gl.Viewport(0, 0, int32(winWidth), int32(winHeight))

	running = true
	for running {
		for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				running = false
			case *sdl.MouseMotionEvent:
				fmt.Printf("[%d ms] MouseMotion\tid:%d\tx:%d\ty:%d\txrel:%d\tyrel:%d\n", t.Timestamp, t.Which, t.X, t.Y, t.XRel, t.YRel)
			}
		}
		drawgl()
		window.GLSwap()
	}
}

func drawgl() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	gl.Begin(gl.TRIANGLES)
	gl.Color3f(1.0, 0.0, 0.0)
	gl.Vertex2f(0.5, 0.0)
	gl.Color3f(0.0, 1.0, 0.0)
	gl.Vertex2f(-0.5, -0.5)
	gl.Color3f(0.0, 0.0, 1.0)
	gl.Vertex2f(-0.5, 0.5)
	gl.End()
}

/*
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
*/
