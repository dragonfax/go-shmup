package main

import (
	"math"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

var player = Player{}

type Player struct {
	X            int32
	Y            int32
	Texture      *sdl.Texture
	Rect         sdl.Rect
	AngleDegrees float64
}

func (p *Player) prepSprite() {
	p.Rect = sdl.Rect{X: 0, Y: 0, W: 10, H: 10}

	pf, err := window.GetPixelFormat()
	if err != nil {
		panic(err)
	}

	t, err := renderer.CreateTexture(pf, sdl.TEXTUREACCESS_TARGET, 10, 10)
	if err != nil {
		panic(err)
	}
	player.Texture = t

	err = renderer.SetRenderTarget(t)
	if err != nil {
		panic(err)
	}
	defer renderer.SetRenderTarget(nil)

	renderer.SetDrawColor(0, 0, 0, 0)
	renderer.Clear()

	renderer.SetDrawColor(255, 0, 0, 255)
	renderer.DrawLine(5, 0, 0, 10)
	renderer.DrawLine(5, 0, 10, 10)
	renderer.Present()
}

func (p *Player) Draw() {
	renderer.CopyEx(player.Texture, &p.Rect, &sdl.Rect{X: p.X, Y: p.Y, W: 10, H: 10}, p.AngleDegrees, nil, sdl.FLIP_NONE)
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
			p.UpdateAngle()
			fire(player.X, player.Y)
		}

		time.Sleep(time.Second / 30)
	}
}

func (p *Player) UpdateAngle() {
	var xv, yv int16
	sdl.Do(func() {
		xv = controller.Axis(sdl.CONTROLLER_AXIS_RIGHTX)
		yv = controller.Axis(sdl.CONTROLLER_AXIS_RIGHTY)
	})
	r := math.Atan2(float64(yv), float64(xv))
	a := radians2Degrees(r)
	a += 90
	p.AngleDegrees = a
}

func radians2Degrees(r float64) (degrees float64) {
	return r * 180 / math.Pi
}
