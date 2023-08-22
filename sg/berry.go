package sg

import (
	"math/rand"

	"github.com/veandco/go-sdl2/sdl"
)

type Berry struct {
	Position Pixel
}

func NewBerry(snake Snake) Berry {
	b := Berry{}

	for {
		b.Position = Pixel{RandomInt(1, GRID_SIZE-2), RandomInt(1, GRID_SIZE-2)}

		if !snake.Contains(b.Position) {
			break
		}
	}

	return b
}

func (b *Berry) Render(r *sdl.Renderer, color sdl.Color) {
	b.Position.Render(r, color)
}

func RandomInt(min, max int) int {
	return rand.Intn(max-min+1) + min
}
