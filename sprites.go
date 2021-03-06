package main

import "github.com/tbruyelle/fsm"

type Direction int

const (
	Up    Direction = 1
	Down  Direction = -1
	Left  Direction = 2
	Right Direction = -2
)

var (
	ratio          float32
	snakeH, snakeW float32
)

func init() {
	ratio = 0.4
	snakeW, snakeH = SnakeW*ratio, SnakeH*ratio
}

type Snake struct {
	fsm.Object
	Dir    Direction
	Size   int
	Speed  float32
	body   *fsm.Object
	tongue *fsm.Object
}

type Queue struct {
	fsm.Object
	pos float32
}

func NewSnake(x, y float32) *Snake {
	s := &Snake{Dir: Left, Size: 0, Speed: 2}
	s.X = x
	s.Y = y
	s.Width = 1
	s.Height = 1
	s.Action = fsm.ActionFunc(snakeMove)
	s.Register(scene, eng)
	// The tongue
	t := &fsm.Object{
		X:      (25 + TongueW) * ratio,
		Y:      114 * ratio,
		Width:  TongueW * ratio,
		Height: TongueH * ratio,
		Sprite: texs[texTongue],
		Action: fsm.ActionFunc(tongueOut),
	}
	t.Register(&s.Object, eng)
	s.tongue = t
	// The body container
	bc := &fsm.Object{
		X:      0,
		Y:      0,
		Width:  1,
		Height: 1,
		Action: fsm.ActionFunc(bodyBounceOut),
	}
	bc.Register(&s.Object, eng)

	// The eye
	e := &fsm.Object{
		X:      49 * ratio,
		Y:      70 * ratio,
		Width:  EyeW * ratio,
		Height: EyeH * ratio,
		Sprite: texs[texEye],
	}
	e.Register(bc, eng)
	// The pupille
	p := &fsm.Object{
		X:      64 * ratio,
		Y:      70 * ratio,
		Width:  PupilleW * ratio,
		Height: PupilleH * ratio,
		Sprite: texs[texPupille],
		Action: fsm.Wait{
			Until: 15,
			Next:  fsm.ActionFunc(pupilleFollow),
		},
	}
	p.Register(bc, eng)

	// The body
	b := &fsm.Object{
		X:      0,
		Y:      0,
		Width:  snakeW,
		Height: snakeH,
		Sprite: texs[texSnakeHead],
	}
	b.Register(bc, eng)
	s.body = b
	return s
}

func (s *Snake) Inc() {
	q := &Queue{
		pos: float32(s.Size),
	}
	q.Sprite = texs[texSnakeQueue]
	q.Width = QueueW * ratio
	q.Height = QueueH * ratio
	q.Data = q
	q.Action = fsm.ActionFunc(queueMove)
	q.Register(scene, eng)
	s.Size++
}

func (s *Snake) Collide(o *fsm.Object) bool {
	if o.X >= s.X+s.body.Width || o.X+o.Width <= s.X || o.Y >= s.Y+s.body.Height || o.Y+o.Height <= s.Y {
		return false
	}
	return true
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
	c.Width = AppleW * ratio
	c.Height = AppleH * ratio
	c.Sprite = texs[texApple]
	c.Register(scene, eng)
	objs = append(objs, &c.Object)
	return c
}

func NewCherry(x, y float32) *Cherry {
	c := &Cherry{}
	c.X = x
	c.Y = y
	c.Width = CherryW * ratio
	c.Height = CherryH * ratio
	c.Sprite = texs[texCherry]
	c.Register(scene, eng)
	objs = append(objs, &c.Object)
	return c
}
