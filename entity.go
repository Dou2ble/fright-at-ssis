package main

import "github.com/hajimehoshi/ebiten/v2"

type Entity interface {
	Draw(*ebiten.Image)
}
