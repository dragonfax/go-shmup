package main

import (
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

	})
}
func drawRect() {

	renderer.SetDrawColor(0, 0, 0, 255)
	renderer.Clear()

	player.Draw()

	drawBullets()

	drawMonsters()

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
