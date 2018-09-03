package main

import (
	"math"
	"sync"
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

		if abs(b.X-player.X) > 1000 || abs(b.Y-player.Y) > 1000 {
			b.remove()
			break
		}

		m := findMonsterCollision(b.X, b.Y)
		if m != nil {
			m.remove()
		}

		time.Sleep(time.Second / 30)
	}
}

func (b *Bullet) remove() {
	b.Dead = true
	bulletsLock.Lock()
	var i int
	for i = 0; i < len(bullets); i++ {
		if bullets[i] == b {
			break
		}
	}
	//remove this bullet
	bullets = append(bullets[:i], bullets[i+1:]...)
	bulletsLock.Unlock()

}

var bullets []*Bullet
var bulletsLock = &sync.Mutex{}

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
	bulletsLock.Lock()
	bullets = append(bullets, bullet)
	bulletsLock.Unlock()
	go bullet.move()
}
