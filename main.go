package main

import (
	"fmt"
	"math"
	"sync"

	"github.com/veandco/go-sdl2/sdl"
)

var player = Player{}

func degreesToRadians(d float64) float64 {
	return d * math.Pi / 180
}

func radiansToDegrees(r float64) float64 {
	return r * 180 / math.Pi
}

var controller *sdl.GameController
var renderer *sdl.Renderer

const DEADZONE = 3200

func abs(x int32) int32 {
	if x > 0 {
		return x
	} else if x < 0 {
		return -x
	} else {
		return 0
	}
}

func abs16(x int16) int16 {
	if x > 0 {
		return x
	} else if x < 0 {
		return -x
	} else {
		return 0
	}
}

var window *sdl.Window

func run() {
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
	// defer sdl.Quit()
	// defer window.Destroy()

	go render()
	go player.moveAndFire()

	done := &sync.WaitGroup{}
	done.Add(1)
	go handleInput(done)

	// wait for input to say we're done.
	done.Wait()
}

func main() {
	sdl.Main(run)
}
