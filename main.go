package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"snake/sg"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

// Game variables
var border sg.Border
var snake sg.Snake
var berry sg.Berry
var direction string
var newDirection string
var gameOver bool = false
var font ttf.Font

//go:embed assets\Arial.ttf
var fontData []byte

func main() {
	setupGame()
	setupWindow()
}

func setupGame() {
	border = sg.NewBorder(sg.GRID_SIZE)
	snake = sg.NewSnake(sg.SNAKE_SIZE)
	berry = sg.NewBerry(snake)

	direction = "right"
	newDirection = direction
	gameOver = false
}

func setupWindow() {
	// Initialize SDL
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Println("An error occured while initializing SDL:", err)
		return
	}
	defer sdl.Quit()

	// Create window
	window, err := sdl.CreateWindow("Snake Go", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, sg.SIZE, sg.SIZE, sdl.WINDOW_OPENGL)
	if err != nil {
		fmt.Println("An error occured while window:", err)
		return
	}
	defer window.Destroy()

	// Create renderer for the window
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println("An error occured while creating renderer:", err)
		return
	}
	defer renderer.Destroy()

	// Initialize SDL TTF
	if err := ttf.Init(); err != nil {
		fmt.Println("An error occured while initializing SDL-TTF:", err)
		return
	}
	defer ttf.Quit()

	// Get RWops from font file data
	rw, err := sdl.RWFromMem(fontData)
	if err != nil {
		fmt.Println("An error occured while loading RWops for font:", err)
		return
	}

	// Load font
	f, err := ttf.OpenFontRW(rw, 1, 16*(sg.SCALE/10))
	if err != nil {
		fmt.Println("An error occured while loading font:", err)
		return
	}
	font = *f

	// Game loop
	var lastTick int64
	for {
		// Check for window events
		for e := sdl.PollEvent(); e != nil; e = sdl.PollEvent() {
			switch e := e.(type) {
			case *sdl.QuitEvent:
				return
			case *sdl.KeyboardEvent:
				if e.State == sdl.RELEASED {
					break
				}
				key := e.Keysym.Scancode
				keyPressed(key)
				break
			}
		}

		// Check if it's time for the next tick
		now := time.Now().UnixMilli()
		if now-lastTick < sg.TICK_TIME {
			continue
		}
		lastTick = now

		// Tick and render the game
		tick()
		render(renderer)
	}
}

func tick() {
	direction = newDirection

	// Move snake and check if it actually moved
	if !snake.Move(direction, border) {
		// Game over
		gameOver = true
		return
	}

	// Check if snake got the berry
	if snake.Contains(berry.Position) {
		berry = sg.NewBerry(snake)
		snake.Grow()
	}
}

func render(r *sdl.Renderer) {
	// Clear background
	c := sg.COLORS["background"]
	r.SetDrawColor(c.R, c.G, c.B, c.A)
	r.Clear()

	if gameOver {
		scale := int(math.Floor(sg.SCALE * 1.5))
		score := snake.GetSize() - sg.SNAKE_SIZE

		drawText(r, "Game over!", sg.SIZE/2, sg.SIZE/2-(scale*2), sg.COLORS["gameOver"])
		drawText(r, fmt.Sprintf("Score: %d", score), sg.SIZE/2, sg.SIZE/2-scale, sg.COLORS["scoreNumber"])
		drawText(r, "Score: "+strings.Repeat(" ", len(strconv.Itoa(score))*2), sg.SIZE/2, sg.SIZE/2-scale, sg.COLORS["score"])

		border.Render(r, sg.COLORS["border"])
		r.Present()
		return
	}

	// Render everything
	border.Render(r, sg.COLORS["border"])
	snake.Render(r, sg.COLORS["snakeHead"], sg.COLORS["snakeBody"])
	berry.Render(r, sg.COLORS["berry"])

	r.Present()
}

func drawText(r *sdl.Renderer, text string, x int, y int, color sdl.Color) {
	surface, err := font.RenderUTF8Solid(text, color)
	if err != nil {
		println("An error occured while rendering text:", err)
		return
	}
	defer surface.Free()

	texture, err := r.CreateTextureFromSurface(surface)
	if err != nil {
		println("Error while CreateTextureFromSurface!")
	}
	defer texture.Destroy()

	rect := sdl.Rect{X: int32(x) - (surface.W / 2), Y: int32(y) - (surface.H / 2), W: surface.W, H: surface.H}
	r.Copy(texture, nil, &rect)
}

func keyPressed(key sdl.Scancode) {
	if gameOver {
		return
	}

	switch key {
	case sdl.SCANCODE_UP, sdl.SCANCODE_W:
		if direction == "down" {
			break
		}
		newDirection = "up"
		break
	case sdl.SCANCODE_DOWN, sdl.SCANCODE_S:
		if direction == "up" {
			break
		}
		newDirection = "down"
		break
	case sdl.SCANCODE_LEFT, sdl.SCANCODE_A:
		if direction == "right" {
			break
		}
		newDirection = "left"
		break
	case sdl.SCANCODE_RIGHT, sdl.SCANCODE_D:
		if direction == "left" {
			break
		}
		newDirection = "right"
		break
	}
}
