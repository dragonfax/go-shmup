package main

import (
	"math"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

type Bullet struct {
	X, Y  int32
	Angle float64 // in Radians
	Dead  bool
}

func (b *Bullet) draw() {
	if !b.Dead {
		renderer.SetDrawColor(0, 255, 0, 255)
		renderer.DrawPoint(b.X, b.Y)
	}
}

func (b *Bullet) move() {
	for {
		b.X += int32(math.Cos(b.Angle) * 10)
		b.Y += int32(math.Sin(b.Angle) * 10)

		if abs(b.X-player.X) > 100 || abs(b.Y-player.Y) > 100 {
			// remove this bullet
			b.Dead = true
			break
		}

		time.Sleep(time.Second / 30)
	}
}

var bullets []*Bullet

func drawBullets() {
	for i := 0; i < len(bullets); i++ {
		bullets[i].draw()
	}
}

func fire(x, y int32) {
	var xv, yv int16
	sdl.Do(func() {
		xv = controller.Axis(sdl.CONTROLLER_AXIS_RIGHTX)
		yv = controller.Axis(sdl.CONTROLLER_AXIS_RIGHTY)
	})
	r := math.Atan2(float64(yv), float64(xv))
	bullet := &Bullet{X: x, Y: y, Angle: r}
	bullets = append(bullets, bullet)
	go bullet.move()
}
