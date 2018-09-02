package main

type Player struct {
	X int32
	Y int32
}

func (p *Player) Draw() {
	renderer.SetDrawColor(255, 0, 0, 255)
	renderer.DrawLine(p.X, p.Y, p.X+10, p.Y+10)
	renderer.DrawLine(p.X, p.Y, p.X-10, p.Y+10)
}
