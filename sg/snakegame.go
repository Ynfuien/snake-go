package sg

import "github.com/veandco/go-sdl2/sdl"

// Configuration variables
const (
	GRID_SIZE  = 32
	SNAKE_SIZE = 5
	SCALE      = 20
	SIZE       = GRID_SIZE*SCALE + GRID_SIZE - 1
	TICK_TIME  = 100
)

// Game colors
var COLORS = map[string]sdl.Color{
	"background":  {R: 36, G: 36, B: 36, A: 255},
	"snakeHead":   {R: 255, G: 170, B: 0, A: 255},
	"snakeBody":   {R: 255, G: 255, B: 85, A: 255},
	"berry":       {R: 255, G: 85, B: 85, A: 255},
	"border":      {R: 85, G: 85, B: 85, A: 255},
	"gameOver":    {R: 255, G: 85, B: 85, A: 255},
	"score":       {R: 255, G: 255, B: 85, A: 255},
	"scoreNumber": {R: 255, G: 170, B: 0, A: 255},
}
