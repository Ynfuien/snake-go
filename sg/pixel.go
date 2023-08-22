package sg

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Pixel struct {
	x int
	y int
}

func (p *Pixel) Render(renderer *sdl.Renderer, c sdl.Color) {
	renderer.SetDrawColor(c.R, c.G, c.B, c.A)
	rect := sdl.Rect{X: int32(p.x)*SCALE + int32(p.x), Y: int32(p.y)*SCALE + int32(p.y), W: SCALE, H: SCALE}
	renderer.FillRect(&rect)
}

func (p *Pixel) Equals(pixel Pixel) bool {
	return p.x == pixel.x && p.y == pixel.y
}
