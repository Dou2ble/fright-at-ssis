package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Player struct {
	x float32
	y float32
}

func (p Player) Draw(image *ebiten.Image) {
	// character
	vector.DrawFilledRect(image, p.x, p.y, PlayerWidth, PlayerHeight, color.Black, false)
}
