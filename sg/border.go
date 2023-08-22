package sg

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Border struct {
	borderPixels []Pixel
}

func NewBorder(size int) Border {
	b := Border{}
	bp := b.borderPixels

	for i := 0; i <= size-1; i++ {
		// Border in width
		bp = append(bp, Pixel{i, 0})
		bp = append(bp, Pixel{i, size - 1})

		// Border in height
		if i == 0 || i == size-1 {
			continue
		}
		bp = append(bp, Pixel{0, i})
		bp = append(bp, Pixel{size - 1, i})
	}

	b.borderPixels = bp
	return b
}

func (b *Border) Render(r *sdl.Renderer, color sdl.Color) {
	for _, p := range b.borderPixels {
		p.Render(r, color)
	}
}

func (b *Border) Contains(pixel Pixel) bool {
	for _, p := range b.borderPixels {
		if p.Equals(pixel) {
			return true
		}
	}

	return false
}
