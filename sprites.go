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
	// 88x80
	AppleW, AppleH = float32(24.6), float32(22.4)
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
	s.Node(scene, eng)
	return s
}

type Cherry struct {
	fsm.Object
}

type Apple struct {
	fsm.Object
}

func NewApple(x, y float32) *Apple {
	c := &Apple{}
	c.X = x
	c.Y = y
	c.Width = AppleW
	c.Height = AppleH
	c.Sprite = texs[texApple]
	c.Node(scene, eng)
	objs = append(objs, &c.Object)
	return c
}

func NewCherry(x, y float32) *Cherry {
	c := &Cherry{}
	c.X = x
	c.Y = y
	c.Width = CherryW
	c.Height = CherryH
	c.Sprite = texs[texCherry]
	c.Node(scene, eng)
	objs = append(objs, &c.Object)
	return c
}
