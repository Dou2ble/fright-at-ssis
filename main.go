package main

import (
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"

	//input "github.com/hajimehoshi/ebiten/v2/inpututil"
	"image/color"

	// "fmt"

	"image"
	_ "image/png"
)

const (
	TotalWidth  = 400
	TotalHeight = 224

	PlayerWidth  = 18
	PlayerHeight = 32

	FloorHeight = 33
)

type Player struct {
	x int
	y int
}

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
	if (ebiten.IsKeyPressed(ebiten.KeyUp) || ebiten.IsKeyPressed(ebiten.KeyW)) && TotalHeight-g.player.y-PlayerHeight <= FloorHeight {
		g.player.y--
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) || ebiten.IsKeyPressed(ebiten.KeyS) {
		g.player.y++
	}


	return nil
}

func DrawGame(g *Game) *ebiten.Image {
	file, _ := os.Open("background.png")
	defer file.Close()
	img, _, _ := image.Decode(file)
	sprite := ebiten.NewImageFromImage(img)


	image := ebiten.NewImage(TotalWidth, TotalHeight)
	image.Fill(color.White)

	//floor
	vector.DrawFilledRect(image, 0, float32(TotalHeight-FloorHeight), float32(TotalWidth), float32(FloorHeight), color.RGBA{255, 0, 0, 255}, false)

	image.DrawImage(sprite, nil)

	// character
	vector.DrawFilledRect(image, float32(g.player.x), float32(g.player.y), PlayerWidth, PlayerHeight, color.Black, false)

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

	game := &Game{Player{10, 200}}
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
