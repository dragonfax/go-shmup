package main

import (
	"math/rand"
	"sync"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

type Monster struct {
	X, Y int32
}

var monsterList []*Monster
var monsterLock = &sync.Mutex{}

func drawMonsters() {
	for _, m := range monsterList {
		rect := &sdl.Rect{X: m.X, Y: m.Y, W: 10, H: 10}
		renderer.SetDrawColor(0, 0, 255, 255)
		renderer.DrawRect(rect)
	}
}

func (m *Monster) move() {
	for {
		if player.X-m.X > 0 {
			m.X += 1
		} else if player.X-m.X < 0 {
			m.X -= 1
		}

		if player.Y-m.Y > 0 {
			m.Y += 1
		} else if player.Y-m.Y < 0 {
			m.Y -= 1
		}

		time.Sleep(time.Second / 10)
	}
}

func monsterSpawner() {
	for {
		addMonster()
		time.Sleep(time.Second)
	}
}

func addMonster() {

	x := rand.Int31n(400)
	y := rand.Int31n(400)

	m := &Monster{X: x, Y: y}
	monsterLock.Lock()
	monsterList = append(monsterList, m)
	monsterLock.Unlock()

	go m.move()
}

func findMonsterCollision(x, y int32) *Monster {
	for _, m := range monsterList {
		if (x >= m.X && x <= m.X+10) && (y >= m.Y && y <= m.X+10) {
			return m
		}
	}
	return nil
}

func (m *Monster) remove() {
	monsterLock.Lock()

	var i int
	for i = 0; i < len(monsterList); i++ {
		if monsterList[i] == m {
			break
		}
	}

	monsterList = append(monsterList[:i], monsterList[i+1:]...)
	monsterLock.Unlock()
}
