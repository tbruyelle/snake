// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/tbruyelle/fsm"
	"image"
	"log"
	"time"

	"image/color"
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
				snake.Action = &snakeTurn{
					dir: Left,
				}
			} else {
				snake.Action = &snakeTurn{
					dir: Right,
				}
			}
		case Left, Right:
			if t.Loc.Y.Px() < snake.Y {
				snake.Action = &snakeTurn{
					dir: Up,
				}
			} else {
				snake.Action = &snakeTurn{
					dir: Down,
				}
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
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

	texs = loadTextures()
	scene = &sprite.Node{}
	eng.Register(scene)
	eng.SetTransform(scene, f32.Affine{
		{1, 0, 0},
		{0, 1, 0},
	})

	var n *sprite.Node

	// Background
	bg := newNode()
	w, h := int(geom.Width.Px()), int(geom.Height.Px())
	texbg, err := fsm.LoadColorTexture(eng, color.RGBA{237, 201, 175, 255}, w, h)
	if err != nil {
		log.Fatal(err)
	}
	eng.SetSubTex(bg, texbg)
	eng.SetTransform(bg, f32.Affine{
		{geom.Width.Px(), 0, 0},
		{0, geom.Height.Px(), 0},
	})

	n = newNode()
	snake = NewSnake(float32(geom.Width/2), float32(geom.Height/2))
	n.Arranger = &snake.Object

	n = newNode()
	eng.SetSubTex(n, texs[texCerise])
	eng.SetTransform(n, f32.Affine{
		{CherryW, 0, 20},
		{0, CherryH, 40},
	})

	// Snake
	n = newNode()
	snake = NewSnake(float32(geom.Width/2), float32(geom.Height/2))
	n.Arranger = snake

}

const (
	texSnakeHead = iota
	texCerise
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
		texSnakeHead: sprite.SubTex{t, image.Rect(0, 0, 280, 184)},
		texCerise:    sprite.SubTex{t, image.Rect(0, 184, 80, 184+80)},
	}
}

type arrangerFunc func(e sprite.Engine, n *sprite.Node, t clock.Time)

func (a arrangerFunc) Arrange(e sprite.Engine, n *sprite.Node, t clock.Time) { a(e, n, t) }
