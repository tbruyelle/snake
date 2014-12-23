package main

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

type Sprite struct {
	X, Y float32
	W, H float32
}

type Snake struct {
	Sprite
	Dir   Direction
	Size  int
	Speed float32
}

func NewSnake(x, y float32) *Snake {
	s := &Snake{Dir: Left, Size: 1, Speed: 2}
	s.X = x
	s.Y = y
	s.W = SnakeW
	s.H = SnakeH
	return s
}

type Cherry struct {
	Sprite
}

func NewCherry(x, y float32) *Cherry {
	c := &Cherry{}
	c.X = x
	c.Y = y
	c.W = CherryW
	c.H = CherryH
	return c
}
