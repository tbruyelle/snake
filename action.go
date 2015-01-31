package main

import (
	"math"

	"github.com/tbruyelle/fsm"
	"golang.org/x/mobile/geom"
	"golang.org/x/mobile/sprite/clock"
)

const (
	halfPi = math.Pi / 2
)

func snakeMove(o *fsm.Object, t clock.Time) {
	switch snake.Dir {
	case Up:
		o.Vx = 0
		o.Vy = -snake.Speed
		o.Rx = o.X + snake.body.Width/2
		o.Ry = o.Y + snake.body.Height/2
		o.Angle = halfPi
	case Left:
		o.Vx = -snake.Speed
		o.Vy = 0
		o.Rx = 0
		o.Ry = 0
		o.Angle = 0
	case Down:
		o.Vx = 0
		o.Vy = snake.Speed
		o.Rx = o.X + snake.body.Width/2
		o.Ry = o.Y + snake.body.Height/2
		o.Angle = -halfPi
	case Right:
		o.Vx = snake.Speed
		o.Vy = 0
		o.Rx = o.X + snake.body.Width/2
		o.Ry = o.Y + snake.body.Height/2
		o.Angle = -math.Pi
	}
	if snake.X > float32(geom.Width) {
		snake.X = -snake.body.Width
	}
	if snake.X+snake.body.Width < 0 {
		snake.X = float32(geom.Width)
	}
	if snake.Y > float32(geom.Height) {
		snake.Y = -snake.body.Height
	}
	if snake.Y+snake.body.Height < 0 {
		snake.Y = float32(geom.Height)
	}
}

type snakeTurn struct {
	dir       Direction
	angle     float32
	origAngle float32
}

func (a *snakeTurn) Do(o *fsm.Object, t clock.Time) {
	if o.Time == 0 {
		o.Time = t
		o.Vx, o.Vy = 0, 0
		o.Rx = o.X + snake.body.Width/2
		o.Ry = o.Y + snake.body.Height/2
		switch {
		case snake.Dir == Up && a.dir == Left,
			snake.Dir == Left && a.dir == Down,
			snake.Dir == Down && a.dir == Right,
			snake.Dir == Right && a.dir == Up:
			a.angle = -halfPi
		case snake.Dir == Up && a.dir == Right,
			snake.Dir == Left && a.dir == Up,
			snake.Dir == Down && a.dir == Left,
			snake.Dir == Right && a.dir == Down:
			a.angle = halfPi
		}
		a.origAngle = o.Angle
	}
	f := clock.EaseOut(o.Time, o.Time+6, t)
	o.Angle = a.origAngle + f*a.angle
	if f == 1 {
		o.Time = 0
		snake.Dir = a.dir
		o.Action = fsm.ActionFunc(snakeMove)
	}

}

func queueMove(o *fsm.Object, t clock.Time) {
	q := o.Data.(*Queue)
	o.X = snake.X + snakeW + (q.pos * o.Width)
	o.Y = snake.Y

	o.Vx = snake.Vx
	o.Vy = snake.Vy
	o.Angle = snake.Angle
	o.Rx = snake.Rx
	o.Ry = snake.Ry
}
