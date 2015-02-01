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
	texTongue
	texEye
	texPupille

	SnakeW, SnakeH     = 280, 184
	QueueW, QueueH     = 138, 138
	CherryW, CherryH   = 84, 84
	AppleW, AppleH     = 102, 88
	TongueW, TongueH   = 57, 31
	EyeW, EyeH         = 41, 26
	PupilleW, PupilleH = 16, 18
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
		texTongue:     sprite.SubTex{t, image.Rect(0, 367, TongueW, 367+TongueH)},
		texEye:        sprite.SubTex{t, image.Rect(0, 403, EyeW, 403+EyeH)},
		texPupille:    sprite.SubTex{t, image.Rect(66, 410, 66+PupilleW, 410+PupilleH)},
	}
}
