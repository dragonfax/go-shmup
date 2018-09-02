package main

import (
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
	// window.UpdateSurface()
	renderer.Present()
}

var surface *sdl.Surface
var window *sdl.Window
var joystick *sdl.Joystick
var renderer *sdl.Renderer

func main() {
	var err error
	if err = sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

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
	if n > 0 {
		sdl.JoystickEventState(sdl.ENABLE)
		joystick = sdl.JoystickOpen(0)
	}

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
			case *sdl.JoyAxisEvent:
				if e.Which == 0 {
					if e.Axis == 0 {
						if e.Value > 0 {
							x += 1
						} else if e.Value < 0 {
							x -= 1
						}
					} else if e.Axis == 1 {
						if e.Value > 0 {
							y += 1
						} else if e.Value < 0 {
							y -= 1
						}
					}
				}
			}
		}
		sdl.Delay(1000 / 60)
	}
}
