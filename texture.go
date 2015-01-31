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
		texApple:      sprite.SubTex{t, image.Rect(0, 273, 102, 273+88)},
	}
}
