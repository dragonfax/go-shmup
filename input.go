package main

import (
	"sync"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

func handleInput(done *sync.WaitGroup) {
	running := true
	for running {
		sdl.Do(func() {
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
				}
			}
		})
		time.Sleep(time.Second / 60)
	}
	done.Done()
}
