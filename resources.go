package main

import (
	"image"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	SpriteBackground ebiten.Image
)

const (
	GFX   = "GFX"
	SFX   = "SFX"
	Music = "music"
)

func resourcesFile(name string, kind string) string {
	return "resources" + string(os.PathSeparator) + kind + string(os.PathSeparator) + name
}

func loadSprite(path string) *ebiten.Image {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("Failed to load"+path, err)
	}

	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatal("Failed to load"+path, err)
	}

	file.Close()

	return ebiten.NewImageFromImage(img)
}

func LoadResources() {
	SpriteBackground = *loadSprite(resourcesFile("background.png", GFX))
}
