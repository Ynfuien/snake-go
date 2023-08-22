package sg

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Snake struct {
	bodyPixels []Pixel
	headPixel  Pixel
}

func NewSnake(size int) Snake {
	s := Snake{}
	bp := s.bodyPixels

	s.headPixel = Pixel{GRID_SIZE/2 + (size / 2), GRID_SIZE/2 - 1}
	for i := size - 1; i > 0; i-- {
		bp = append(bp, Pixel{s.headPixel.x - i, s.headPixel.y})
	}

	s.bodyPixels = bp
	return s
}

func (s *Snake) Render(r *sdl.Renderer, headColor sdl.Color, bodyColor sdl.Color) {
	s.headPixel.Render(r, headColor)
	for _, p := range s.bodyPixels {
		p.Render(r, bodyColor)
	}
}

func (s *Snake) Move(direction string, border Border) bool {
	x := s.headPixel.x
	y := s.headPixel.y

	if direction == "up" {
		y--
	} else if direction == "right" {
		x++
	} else if direction == "down" {
		y++
	} else if direction == "left" {
		x--
	}

	newHead := Pixel{x, y}
	if s.Contains(newHead) {
		return false
	}
	if border.Contains(newHead) {
		return false
	}

	s.bodyPixels = append(s.bodyPixels, s.headPixel)
	s.bodyPixels = s.bodyPixels[1:]
	s.headPixel = newHead
	return true
}

func RemoveIndex[T any](slice []T, index int) []T {
	return append(slice[:index], slice[index+1:]...)
}

func (s *Snake) Grow() {
	newBody := Pixel{s.bodyPixels[0].x, s.bodyPixels[0].y}
	s.bodyPixels = append([]Pixel{newBody}, s.bodyPixels...)
}

func (s *Snake) GetSize() int {
	return len(s.bodyPixels) + 1
}

func (s *Snake) Contains(pixel Pixel) bool {
	if s.headPixel.Equals(pixel) {
		return true
	}

	for _, p := range s.bodyPixels {
		if p.Equals(pixel) {
			return true
		}
	}

	return false
}
