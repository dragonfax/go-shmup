package main

import (
	"fmt"
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

func drawRect(x, y int32) {
	// surface.FillRect(nil, 0)
	// rect := sdl.Rect{X: x, Y: y, W: 200, H: 200}
	// surface.FillRect(&rect, 0xffff0000)

	renderer.SetDrawColor(0, 0, 0, 255)
	renderer.Clear()

	renderer.SetDrawColor(255, 0, 0, 255)
	renderer.DrawLine(x, y, x+10, y+10)
	renderer.DrawLine(x, y, x-10, y+10)

	for i := 0; i < len(bullets); i++ {
		b := bullets[i]
		renderer.DrawPoint(b.X, b.Y)
	}

	// window.UpdateSurface()
	renderer.Present()
}

type Bullet struct {
	X, Y  int32
	Angle float64 // in Radians
}

var bullets []Bullet

func fire(x, y int32) {
	xv := controller.Axis(sdl.CONTROLLER_AXIS_RIGHTX)
	yv := controller.Axis(sdl.CONTROLLER_AXIS_RIGHTY)
	r := math.Atan2(float64(yv), float64(xv))
	fmt.Println("fire angle ", r)
	bullets = append(bullets, Bullet{X: x, Y: y, Angle: r})
}

func degreesToRadians(d float64) float64 {
	return d * math.Pi / 180
}

func radiansToDegrees(r float64) float64 {
	return r * 180 / math.Pi
}

func moveBullets() {
	for i := 0; i < len(bullets); i++ {
		b := bullets[i]
		b.X += int32(math.Cos(b.Angle) * 10)
		b.Y += int32(math.Sin(b.Angle) * 10)
		bullets[i] = b
	}
}

var controller *sdl.GameController
var renderer *sdl.Renderer

func main() {
	var err error
	if err = sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	var window *sdl.Window
	window, err = sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		800, 600, 0)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	renderer, err = sdl.CreateRenderer(window, -1, 0)
	if err != nil {
		panic(err)
	}

	n := sdl.NumJoysticks()
	// var joystick *sdl.Joystick
	if n < 1 {
		panic(fmt.Sprintf("not enough joysticks %d", n))
	}
	if !sdl.IsGameController(0) {
		panic("no game controller")
	}
	sdl.GameControllerEventState(sdl.ENABLE)
	controller = sdl.GameControllerOpen(0)

	/* surface, err = window.GetSurface()
	if err != nil {
		panic(err)
	} */

	var x, y int32 = 0, 0

	running := true
	for running {
		// fmt.Printf("%d %d\n", x, y)
		drawRect(x, y)
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch e := event.(type) {
			case *sdl.QuitEvent:
				println("Quit")
				running = false
				break
			case *sdl.KeyboardEvent:
				if e.Type == sdl.KEYDOWN {
					if e.Keysym.Sym == sdl.K_q {
						running = false
						break
					}
				}
			case *sdl.ControllerAxisEvent:
				if e.Axis == sdl.CONTROLLER_AXIS_LEFTX {
					if e.Value > 0 {
						x += 1
					} else if e.Value < 0 {
						x -= 1
					}
				} else if e.Axis == sdl.CONTROLLER_AXIS_LEFTY {
					if e.Value > 0 {
						y += 1
					} else if e.Value < 0 {
						y -= 1
					}
				} else if e.Axis == sdl.CONTROLLER_AXIS_RIGHTX || e.Axis == sdl.CONTROLLER_AXIS_RIGHTY {
					fire(x, y)
				}
			}
		}
		moveBullets()
		sdl.Delay(1000 / 60)
	}
}
