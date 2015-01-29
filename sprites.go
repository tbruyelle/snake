package main

import "github.com/tbruyelle/fsm"

type Direction int

const (
	Up    Direction = 1
	Down  Direction = -1
	Left  Direction = 2
	Right Direction = -2

	// ratio 0.28
	// 280x184
	SnakeW, SnakeH = float32(78.4), float32(51.52)
	// 80x80
	CherryW, CherryH = float32(22.4), float32(22.4)
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
	s.Sprite = texs[texSnakeHead]
	s.Action = fsm.ActionFunc(snakeMove)
	return s
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
