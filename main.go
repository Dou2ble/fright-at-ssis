package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"

	//input "github.com/hajimehoshi/ebiten/v2/inpututil"
	"image/color"

	// "fmt"

	_ "image/png"
)

const (
	TotalWidth  = 400
	TotalHeight = 224

	PlayerWidth  = 18
	PlayerHeight = 32

	FloorHeight        = 33
	VirtualFloorHeight = FloorHeight - 5
)

type Game struct {
	player Player
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyLeft) || ebiten.IsKeyPressed(ebiten.KeyA) {
		g.player.x--
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) || ebiten.IsKeyPressed(ebiten.KeyD) {
		g.player.x++
	}
	if (ebiten.IsKeyPressed(ebiten.KeyUp) || ebiten.IsKeyPressed(ebiten.KeyW)) && TotalHeight-g.player.y-PlayerHeight <= VirtualFloorHeight {
		g.player.y--
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) || ebiten.IsKeyPressed(ebiten.KeyS) {
		g.player.y++
	}

	return nil
}

func DrawGame(g *Game) *ebiten.Image {
	image := ebiten.NewImage(TotalWidth, TotalHeight)
	image.Fill(color.White)

	//floor
	vector.DrawFilledRect(image, 0, float32(TotalHeight-FloorHeight), float32(TotalWidth), float32(FloorHeight), color.RGBA{255, 0, 0, 255}, false)

	image.DrawImage(&SpriteBackground, nil)

	g.player.Draw(image)

	return image
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(DrawGame(g), nil)
}

func (g *Game) Layout(_, _ int) (screenWidth, screenHeight int) {
	return TotalWidth, TotalHeight
}

func main() {
	ebiten.SetWindowSize(1280, 720)
	ebiten.SetWindowTitle("SSIS GAME")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	LoadResources()

	game := &Game{Player{10, 200}}
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
