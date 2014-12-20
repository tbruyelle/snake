package main

type Direction int

const (
	Up    Direction = 1
	Down  Direction = -1
	Left  Direction = 2
	Right Direction = -2
)

type Snake struct {
	X, Y  float32
	Dir   Direction
	Size  int
	Speed float32
}

func NewSnake(x, y float32) *Snake {
	s := &Snake{X: x, Y: y, Dir: Left, Size: 1, Speed: 1}
	return s
}
