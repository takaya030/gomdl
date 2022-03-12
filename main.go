package main

import (
	"fmt"
	"io/ioutil"
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/veandco/go-sdl2/sdl"

	"github.com/takaya030/gomdl/mdl"
)


var mdlmodel *mdl.MdlModel
var prev_tick float32
var transx, transy, transz, rotx, roty float32 = 0.0, 0.0, -2.0, 235.0, -90.0

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

	mdlvwInit(`./asset/gsg9.mdl`)
	initGL()
	resizeWindow(winWidth, winHeight)

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

func mdlvwInit(mdl_file string) {

	buf, rferr := ioutil.ReadFile(mdl_file)
	if rferr != nil {
		fmt.Print(rferr)
		return
	}

	// read mdldata
	mdd := mdl.NewMdlData(buf)
	mdlmodel = mdl.NewMdlModel(mdd)
	mdlmodel.Init()
}

func mdlvwDisplay() {
	mdlmodel.SetBlending(0, 0.0)
	mdlmodel.SetBlending(1, 0.0)

	curr_tick := float32(sdl.GetTicks()) / 1000.0
	mdlmodel.AdvanceFrame(curr_tick - prev_tick)
	prev_tick = curr_tick

	mdlmodel.DrawModel()
}

func initGL() {

	// Enable smooth shading
	gl.ShadeModel( gl.SMOOTH )

	// Set the background black
	gl.ClearColor( 0.2, 0.2, 0.2, 0.0 )

	// Depth buffer setup
	gl.ClearDepth( 1.0 )

	// Enables Depth Testing
	gl.Enable( gl.DEPTH_TEST )

	// The Type Of Depth Test To Do
	gl.DepthFunc( gl.LEQUAL )

	// Really Nice Perspective Calculations
	gl.Hint( gl.PERSPECTIVE_CORRECTION_HINT, gl.NICEST )
}

func resizeWindow(width int32, height int32) {

    /* Protect against a divide by zero */
    if  height == 0 {
		height = 1
	}

    /* Height / width ration */
    ratio := float32(width) / float32(height)

    /* Setup our viewport. */
    gl.Viewport( 0, 0, int32(width), int32(height) )

    /* change to the projection matrix and set our viewing volume. */
    gl.MatrixMode( gl.PROJECTION )
    gl.LoadIdentity()

    /* Set our perspective */
    projection := mgl32.Perspective( 45.0, ratio, 0.1, 100.0 )
	gl.LoadMatrixf(&projection[0])

    /* Make sure we're chaning the model view and not the projection */
    gl.MatrixMode( gl.MODELVIEW )

    /* Reset The View */
    gl.LoadIdentity()
}

func drawgl() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	/* Move Left 1.5 Units And Into The Screen 6.0 */
    gl.LoadIdentity()
    gl.Translatef( -1.5, 0.0, -6.0 )

    gl.Begin( gl.TRIANGLES )          /* Drawing Using Triangles       */
      gl.Color3f(   1.0,  0.0,  0.0 ) /* Red                           */
      gl.Vertex3f(  0.0,  1.0,  0.0 ) /* Top Of Triangle               */
      gl.Color3f(   0.0,  1.0,  0.0 ) /* Green                         */
      gl.Vertex3f( -1.0, -1.0,  0.0 ) /* Left Of Triangle              */
      gl.Color3f(   0.0,  0.0,  1.0 ) /* Blue                          */
      gl.Vertex3f(  1.0, -1.0,  0.0 ) /* Right Of Triangle             */
    gl.End()                          /* Finished Drawing The Triangle */

    /* Move Right 3 Units */
    gl.Translatef( 3.0, 0.0, 0.0 )

    /* Set The Color To Blue One Time Only */
    gl.Color3f( 0.5, 0.5, 1.0)

    gl.Begin( gl.QUADS );             /* Draw A Quad              */
      gl.Vertex3f(  1.0,  1.0,  0.0 ) /* Top Right Of The Quad    */
      gl.Vertex3f( -1.0,  1.0,  0.0 ) /* Top Left Of The Quad     */
      gl.Vertex3f( -1.0, -1.0,  0.0 ) /* Bottom Left Of The Quad  */
      gl.Vertex3f(  1.0, -1.0,  0.0 ) /* Bottom Right Of The Quad */
    gl.End()                          /* Done Drawing The Quad    */

	// draw StudioModel
    gl.PushMatrix()
	gl.LoadIdentity()
    gl.Translatef(transx, transy, transz)
	gl.Rotatef(rotx, 0.0, 1.0, 0.0)
    gl.Rotatef(roty, 1.0, 0.0, 0.0)
    gl.Scalef( 0.02, 0.02, 0.02 )
	gl.CullFace( gl.FRONT )
	//gl.Enable( gl.DEPTH_TEST )
    gl.Enable(gl.TEXTURE_2D)

	mdlvwDisplay()

    gl.Disable(gl.TEXTURE_2D)
    gl.PopMatrix()
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
