package main

import (
	"fmt"
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

var player = Player{}

func drawRect(x, y int32) {
	// surface.FillRect(nil, 0)
	// rect := sdl.Rect{X: x, Y: y, W: 200, H: 200}
	// surface.FillRect(&rect, 0xffff0000)

	renderer.SetDrawColor(0, 0, 0, 255)
	renderer.Clear()

	player.Draw()

	for i := 0; i < len(bullets); i++ {
		bullets[i].draw()
	}

	// window.UpdateSurface()
	renderer.Present()
}

func degreesToRadians(d float64) float64 {
	return d * math.Pi / 180
}

func radiansToDegrees(r float64) float64 {
	return r * 180 / math.Pi
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

	running := true
	for running {
		// fmt.Printf("%d %d\n", x, y)
		drawRect(player.X, player.Y)
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
						player.X += 1
					} else if e.Value < 0 {
						player.X -= 1
					}
				} else if e.Axis == sdl.CONTROLLER_AXIS_LEFTY {
					if e.Value > 0 {
						player.Y += 1
					} else if e.Value < 0 {
						player.Y -= 1
					}
				} else if e.Axis == sdl.CONTROLLER_AXIS_RIGHTX || e.Axis == sdl.CONTROLLER_AXIS_RIGHTY {
					fire(player.X, player.Y)
				}
			}
		}
		sdl.Delay(1000 / 60)
	}
}
