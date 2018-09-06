package main

import (
	"sync"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const DEADZONE = 3200

var controller *sdl.GameController

var keyboardMoveUp = false
var keyboardMoveDown = false
var keyboardMoveLeft = false
var keyboardMoveRight = false

func InitJoystick() {
	sdl.Do(func() {
		n := sdl.NumJoysticks()
		if n > 0 {
			if !sdl.IsGameController(0) {
				panic("not a game controller")
			}
			sdl.GameControllerEventState(sdl.ENABLE) // not used
			controller = sdl.GameControllerOpen(0)
		}
	})
}

func handleInput(done *sync.WaitGroup) {
	running := true
	for running {
		sdl.Do(func() {
		INPUT:
			for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
				switch e := event.(type) {
				case *sdl.QuitEvent:
					println("Quit")
					running = false
					break
				case *sdl.KeyboardEvent:
					if e.Type == sdl.KEYDOWN {
						switch e.Keysym.Sym {
						case sdl.K_q:
							running = false
							break INPUT
						case sdl.K_UP:
							keyboardMoveUp = true
						case sdl.K_DOWN:
							keyboardMoveDown = true
						case sdl.K_LEFT:
							keyboardMoveLeft = true
						case sdl.K_RIGHT:
							keyboardMoveRight = true
						}
					} else if e.Type == sdl.KEYUP {
						switch e.Keysym.Sym {
						case sdl.K_UP:
							keyboardMoveUp = false
						case sdl.K_DOWN:
							keyboardMoveDown = false
						case sdl.K_LEFT:
							keyboardMoveLeft = false
						case sdl.K_RIGHT:
							keyboardMoveRight = false
						}
					}
				}
			}
		})
		time.Sleep(time.Second / 60)
	}
	done.Done()
}
