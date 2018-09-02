package main

import (
	"fmt"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

var renderer *sdl.Renderer
var window *sdl.Window

func initRender() {
	sdl.Do(func() {

		var err error
		if err = sdl.Init(sdl.INIT_EVERYTHING); err != nil {
			panic(err)
		}

		window, err = sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
			800, 600, 0)
		if err != nil {
			panic(err)
		}

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
		sdl.GameControllerEventState(sdl.ENABLE) // not used
		controller = sdl.GameControllerOpen(0)

		/* surface, err = window.GetSurface()
		if err != nil {
			panic(err)
		} */

	})
}

func drawRect() {
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

func render() {
	for {
		sdl.Do(func() {
			drawRect()
		})
		time.Sleep(time.Second / 60)
	}
}
