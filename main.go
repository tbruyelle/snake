// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"image"
	"log"
	"time"

	_ "image/png"

	"golang.org/x/mobile/app"
	"golang.org/x/mobile/app/debug"
	"golang.org/x/mobile/event"
	"golang.org/x/mobile/f32"
	"golang.org/x/mobile/geom"
	"golang.org/x/mobile/gl"
	"golang.org/x/mobile/sprite"
	"golang.org/x/mobile/sprite/clock"
	"golang.org/x/mobile/sprite/glsprite"
)

var (
	start     = time.Now()
	lastClock = clock.Time(-1)

	eng   = glsprite.Engine()
	scene *sprite.Node
	texs  []sprite.SubTex
	snake *Snake
)

func main() {
	app.Run(app.Callbacks{
		Start: loadScene,
		Draw:  draw,
		Touch: touch,
	})
}

func draw() {
	// Keep until golang.org/x/mogile/x11.go handle Start callback
	if scene == nil {
		loadScene()
	}

	now := clock.Time(time.Since(start) * 60 / time.Second)
	if now == lastClock {
		// TODO: figure out how to limit draw callbacks to 60Hz instead of
		// burning the CPU as fast as possible.
		// TODO: (relatedly??) sync to vblank?
		return
	}
	lastClock = now

	gl.ClearColor(1, 1, 1, 1)
	gl.Clear(gl.COLOR_BUFFER_BIT)
	eng.Render(scene, now)
	debug.DrawFPS()
}

func touch(t event.Touch) {
	if t.Type == event.TouchEnd {
		switch snake.Dir {
		case Up, Down:
			if t.Loc.X.Px() < snake.X {
				snake.Dir = Left
			} else {
				snake.Dir = Right
			}
		case Left, Right:
			if t.Loc.Y.Px() < snake.Y {
				snake.Dir = Up
			} else {
				snake.Dir = Down
			}
		}
	}
}

func newNode() *sprite.Node {
	n := &sprite.Node{}
	eng.Register(n)
	scene.AppendChild(n)
	return n
}

func loadScene() {
	texs = loadTextures()
	scene = &sprite.Node{}
	eng.Register(scene)
	eng.SetTransform(scene, f32.Affine{
		{1, 0, 0},
		{0, 1, 0},
	})

	var n *sprite.Node

	n = newNode()
	snake = NewSnake(float32(geom.Width/2), float32(geom.Height/2))
	n.Arranger = snake
}

func (s *Snake) Arrange(e sprite.Engine, n *sprite.Node, t clock.Time) {
	snakew, snakeh := float32(72), float32(46)
	var w, h float32
	switch s.Dir {
	case Up:
		s.Y -= s.Speed
		eng.SetSubTex(n, texs[texSnakeHeadU])
		w, h = snakeh, snakew
	case Left:
		s.X -= s.Speed
		eng.SetSubTex(n, texs[texSnakeHeadL])
		w, h = snakew, snakeh
	case Down:
		s.Y += s.Speed
		eng.SetSubTex(n, texs[texSnakeHeadD])
		w, h = snakeh, snakew
	case Right:
		s.X += s.Speed
		eng.SetSubTex(n, texs[texSnakeHeadR])
		w, h = snakew, snakeh
	}
	if s.X-w > geom.Width.Px() {
		s.X = -w
	}
	if s.X+w < 0 {
		s.X = geom.Width.Px()
	}
	if s.Y-h > geom.Height.Px() {
		s.Y = -h
	}
	if s.Y+h < 0 {
		s.Y = geom.Height.Px()
	}

	eng.SetTransform(n, f32.Affine{
		{w, 0, s.X},
		{0, h, s.Y},
	})
}

const (
	texSnakeHeadR = iota
	texSnakeHeadL
	texSnakeHeadU
	texSnakeHeadD
)

func loadTextures() []sprite.SubTex {
	f, err := app.Open("tiles.png")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	img, _, err := image.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	t, err := eng.LoadTexture(img)
	if err != nil {
		log.Fatal(err)
	}

	return []sprite.SubTex{
		texSnakeHeadR: sprite.SubTex{t, image.Rect(0, 0, 256, 164)},
		texSnakeHeadL: sprite.SubTex{t, image.Rect(256, 0, 512, 164)},
		texSnakeHeadU: sprite.SubTex{t, image.Rect(0, 164, 164, 420)},
		texSnakeHeadD: sprite.SubTex{t, image.Rect(164, 164, 328, 420)},
	}
}

type arrangerFunc func(e sprite.Engine, n *sprite.Node, t clock.Time)

func (a arrangerFunc) Arrange(e sprite.Engine, n *sprite.Node, t clock.Time) { a(e, n, t) }
