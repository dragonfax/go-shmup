package main

import (
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

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
