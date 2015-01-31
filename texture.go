package main

import (
	"image"
	"log"

	"golang.org/x/mobile/app"
	"golang.org/x/mobile/sprite"
)

const (
	texSnakeHead = iota
	texSnakeQueue
	texCherry
	texApple

	SnakeW, SnakeH   = 280, 184
	QueueW, QueueH   = 138, 138
	CherryW, CherryH = 84, 84
	AppleW, AppleH   = 102, 88
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
		texSnakeHead:  sprite.SubTex{t, image.Rect(0, 0, SnakeW, SnakeH)},
		texSnakeQueue: sprite.SubTex{t, image.Rect(278, 0, 278+QueueW, QueueH)},
		texCherry:     sprite.SubTex{t, image.Rect(0, 189, CherryW, 189+CherryH)},
		texApple:      sprite.SubTex{t, image.Rect(0, 273, AppleW, 273+AppleH)},
	}
}
