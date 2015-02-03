// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"log"
	"time"

	"github.com/tbruyelle/fsm"

	"image/color"
	_ "image/png"

	"golang.org/x/mobile/app"
	"golang.org/x/mobile/app/debug"
	"golang.org/x/mobile/event"
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
	scene *fsm.Object
	texs  []sprite.SubTex
	snake *Snake
	objs  Objs
	apple *Apple
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
			snake.Inc()
			objs = objs.Remove(i)
			break // one collision per loop?
		}
	}

	eng.Render(scene.Node, now)
	debug.DrawFPS()
}

func touch(t event.Touch) {
	if t.Type == event.TouchEnd {
		switch snake.Dir {
		case Up, Down:
			if float32(t.Loc.X) < snake.X {
				snake.Action = &snakeTurn{
					dir: Left,
				}
			} else {
				snake.Action = &snakeTurn{
					dir: Right,
				}
			}
		case Left, Right:
			if float32(t.Loc.Y) < snake.Y {
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

func loadScene() {
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

	texs = loadTextures()
	scene = &fsm.Object{Width: 1, Height: 1}
	scene.Register(nil, eng)

	// Background
	bg := &fsm.Object{
		Width:  float32(geom.Width),
		Height: float32(geom.Height),
	}
	w, h := int(geom.Width), int(geom.Height)

	texbg, err := fsm.LoadColorTexture(eng, color.RGBA{237, 201, 175, 255}, w, h)
	if err != nil {
		log.Fatal(err)
	}
	bg.Sprite = texbg
	bg.Register(scene, eng)

	objs = make(Objs, 0)

	// a cherry
	NewCherry(20, 40)
	apple = NewApple(150, 50)

	// Snake
	snake = NewSnake(float32(geom.Width/2), float32(geom.Height/2))
}
