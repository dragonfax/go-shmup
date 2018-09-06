package main

import (
	"math"
	"time"

	"github.com/veandco/go-sdl2/gfx"
	"github.com/veandco/go-sdl2/sdl"
)

var player = Player{}

type Player struct {
	X            int32
	Y            int32
	Texture      *sdl.Texture
	AngleDegrees float64
}

const PLAYER_SPRITE_WIDTH = 20 // on screen size.

type Point struct {
	X int32
	Y int32
}

var player_color = sdl.Color{255, 255, 255, 255}
var background_color = sdl.Color{0, 0, 0, 255}

var player_v = []Point{
	Point{
		-PLAYER_SPRITE_WIDTH / 2,
		PLAYER_SPRITE_WIDTH,
	},
	Point{
		0,
		PLAYER_SPRITE_WIDTH * 2,
	},
	Point{
		PLAYER_SPRITE_WIDTH / 2,
		PLAYER_SPRITE_WIDTH,
	},
}

func (player *Player) Draw() {
	pc := Point{ // player center
		X: player.X + PLAYER_SPRITE_WIDTH/2,
		Y: player.Y + PLAYER_SPRITE_WIDTH/2,
	}

	gfx.FilledCircleColor(renderer, pc.X, pc.Y, PLAYER_SPRITE_WIDTH/2, player_color)
	gfx.FilledCircleColor(renderer, pc.X, pc.Y, PLAYER_SPRITE_WIDTH/4, background_color)

	len_p := len(player_v)
	var vx = make([]int16, len_p)
	var vy = make([]int16, len_p)
	for i, p := range player_v {
		pt := Point{p.X + pc.X, p.Y + pc.Y}
		pr := rotate_point(pc, degrees2Radians(player.AngleDegrees+180), pt)
		vx[i] = int16(pr.X)
		vy[i] = int16(pr.Y)
	}
	gfx.PolygonColor(renderer, vx, vy, player_color)
}

// rotate x,y around cx,cy by angle (radians)
func rotate_point(center Point, angle float64, point Point) Point {
	var s = math.Sin(angle)
	var c = math.Cos(angle)

	// translate point back to origin:
	point.X -= center.X
	point.Y -= center.Y

	var xf = float64(point.X)
	var yf = float64(point.Y)

	// rotate point
	xnew := xf*c - yf*s
	ynew := xf*s + yf*c

	// translate point back:
	x := int32(xnew) + center.X
	y := int32(ynew) + center.Y

	return Point{x, y}
}

func (p *Player) moveAndFire() {

	for {
		var lx, ly, rx, ry int16

		if controller != nil {
			sdl.Do(func() {
				lx = controller.Axis(sdl.CONTROLLER_AXIS_LEFTX)
				ly = controller.Axis(sdl.CONTROLLER_AXIS_LEFTY)
				rx = controller.Axis(sdl.CONTROLLER_AXIS_RIGHTX)
				ry = controller.Axis(sdl.CONTROLLER_AXIS_RIGHTY)
			})
		}

		if keyboardMoveUp {
			p.Y -= 1
		}

		if keyboardMoveDown {
			p.Y += 1
		}

		if keyboardMoveLeft {
			p.X -= 1
		}

		if keyboardMoveRight {
			p.X += 1
		}

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

		if abs16(rx) > DEADZONE || abs16(ry) > DEADZONE || keyboardMoveDown || keyboardMoveLeft || keyboardMoveRight || keyboardMoveUp {
			p.UpdateAngle()
			fire(player.X, player.Y)
		}

		time.Sleep(time.Second / 30)
	}
}

func (p *Player) UpdateAngle() {
	var xv, yv int16
	if controller != nil {
		sdl.Do(func() {
			xv = controller.Axis(sdl.CONTROLLER_AXIS_RIGHTX)
			yv = controller.Axis(sdl.CONTROLLER_AXIS_RIGHTY)
		})
	}
	if keyboardMoveDown || keyboardMoveLeft || keyboardMoveRight || keyboardMoveUp {
		yv = 0
		xv = 0
		if keyboardMoveDown {
			yv = 1
		}
		if keyboardMoveLeft {
			xv = -1
		}
		if keyboardMoveRight {
			xv = 1
		}
		if keyboardMoveUp {
			yv = -1
		}
	}
	r := math.Atan2(float64(yv), float64(xv))
	a := radians2Degrees(r)
	a += 90
	p.AngleDegrees = a
}

func radians2Degrees(r float64) (degrees float64) {
	return r * 180 / math.Pi
}

func degrees2Radians(d float64) (radians float64) {
	return d * math.Pi / 180
}
