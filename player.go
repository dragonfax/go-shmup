package main

import (
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

var player = Player{}

type Player struct {
	X int32
	Y int32
}

func (p *Player) Draw() {
	renderer.SetDrawColor(255, 0, 0, 255)
	renderer.DrawLine(p.X, p.Y, p.X+10, p.Y+10)
	renderer.DrawLine(p.X, p.Y, p.X-10, p.Y+10)
}

func (p *Player) moveAndFire() {

	for {
		var lx, ly, rx, ry int16

		sdl.Do(func() {
			lx = controller.Axis(sdl.CONTROLLER_AXIS_LEFTX)
			ly = controller.Axis(sdl.CONTROLLER_AXIS_LEFTY)
			rx = controller.Axis(sdl.CONTROLLER_AXIS_RIGHTX)
			ry = controller.Axis(sdl.CONTROLLER_AXIS_RIGHTY)
		})

		if abs16(lx) > DEADZONE {
			if lx > 0 {
				p.X += 1
			} else if lx < 0 {
				p.X -= 1
			}
		}

		if abs16(ly) > DEADZONE {
			if ly > 0 {
				p.Y += 1
			} else if ly < 0 {
				p.Y -= 1
			}
		}

		if abs16(rx) > DEADZONE || abs16(ry) > DEADZONE {
			fire(player.X, player.Y)
		}

		time.Sleep(time.Second / 30)
	}
}
