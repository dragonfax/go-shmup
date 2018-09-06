package main

import (
	"sync"

	"github.com/veandco/go-sdl2/sdl"
)

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

func run() {
	initRender()
	InitJoystick()

	player.prepSprite()

	go render()
	go player.moveAndFire()

	done := &sync.WaitGroup{}
	done.Add(1)
	go handleInput(done)

	go monsterSpawner()

	// wait for input to say we're done.
	done.Wait()
}

func main() {
	sdl.Main(run)
}
