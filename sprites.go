package main

import (
	"github.com/tbruyelle/fsm"
	"golang.org/x/mobile/geom"
	"golang.org/x/mobile/sprite/clock"
)

type Direction int

const (
	Up    Direction = 1
	Down  Direction = -1
	Left  Direction = 2
	Right Direction = -2

	// ratio 0.28
	// 256x164
	SnakeW, SnakeH = float32(72), float32(46)
	// 59x64
	CherryW, CherryH = float32(16), float32(18)
)

type Snake struct {
	fsm.Object
	Dir   Direction
	Size  int
	Speed float32
}

func NewSnake(x, y float32) *Snake {
	s := &Snake{Dir: Left, Size: 1, Speed: 2}
	s.X = x
	s.Y = y
	s.Width = SnakeW
	s.Height = SnakeH
	s.Action = fsm.ActionFunc(snakeMove)
	return s
}

func snakeMove(o *fsm.Object, t clock.Time) {
	switch snake.Dir {
	case Up:
		o.Vx = 0
		o.Vy = -snake.Speed
		o.Sprite = texs[texSnakeHeadU]
	case Left:
		o.Vx = -snake.Speed
		o.Vy = 0
		o.Sprite = texs[texSnakeHeadL]
	case Down:
		o.Vx = 0
		o.Vy = snake.Speed
		o.Sprite = texs[texSnakeHeadD]
	case Right:
		o.Vx = snake.Speed
		o.Vy = 0
		o.Sprite = texs[texSnakeHeadR]
	}
	if snake.X > float32(geom.Width) {
		snake.X = -snake.Width
	}
	if snake.X+snake.Width < 0 {
		snake.X = float32(geom.Width)
	}
	if snake.Y > float32(geom.Height) {
		snake.Y = -snake.Height
	}
	if snake.Y+snake.Height < 0 {
		snake.Y = float32(geom.Height)
	}
}

type Cherry struct {
	fsm.Object
}

func NewCherry(x, y float32) *Cherry {
	c := &Cherry{}
	c.X = x
	c.Y = y
	c.Width = CherryW
	c.Height = CherryH
	return c
}
