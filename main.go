// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"image"
	"log"
	"time"

	"github.com/tbruyelle/fsm"

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

type Objs []*fsm.Object

func (a Objs) Remove(i int) Objs {
	a[i], a[len(a)-1], a = a[len(a)-1], nil, a[:len(a)-1]
	return a
}

var (
	start     = time.Now()
	lastClock = clock.Time(-1)

	eng   = glsprite.Engine()
	scene *sprite.Node
	texs  []sprite.SubTex
	snake *Snake
	objs  Objs
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

	// test collisions
	for i, o := range objs {
		if snake.Collide(o) {
			o.Dead = true
			snake.Speed++
			snake.Size++
			objs = objs.Remove(i)
			break // one collision per loop?
		}
	}

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

	// Background
	bg := &fsm.Object{Width: float32(geom.Width), Height: float32(geom.Height)}
	w, h := int(geom.Width), int(geom.Height)

	texbg, err := fsm.LoadColorTexture(eng, color.RGBA{237, 201, 175, 255}, w, h)
	if err != nil {
		log.Fatal(err)
	}
	bg.Sprite = texbg
	bg.Node(scene, eng)

	objs = make(Objs, 0)

	// a cherry
	NewCherry(20, 40)
	NewApple(100, 40)

	// Snake
	snake = NewSnake(float32(geom.Width/2), float32(geom.Height/2))
}

const (
	texSnakeHead = iota
	texSnakeQueue
	texCherry
	texApple
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
		texSnakeHead:  sprite.SubTex{t, image.Rect(0, 0, 280, 184)},
		texSnakeQueue: sprite.SubTex{t, image.Rect(278, 0, 278+146, 141)},
		texCherry:     sprite.SubTex{t, image.Rect(0, 184, 80, 184+80)},
		texApple:      sprite.SubTex{t, image.Rect(0, 263, 88, 263+80)},
	}
}
