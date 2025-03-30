package main

import (
	"embed"
	"log"
	"m3g4p0p/game/util"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

//go:embed assets/*
var assets embed.FS
var playerSprite = util.Must(util.LoadImage(assets, "assets/playerShip1_blue.png"))
var fireSprite = util.Must(util.LoadImage(assets, "assets/Effects/fire09.png"))

type Game struct {
	targetPos ebiten.GeoM
}

func (g *Game) Update() error {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		geom := &ebiten.GeoM{}
		geom.Translate(float64(x), float64(y))
		g.targetPos = *geom
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, g.targetPos.String())
	screen.DrawImage(playerSprite, &ebiten.DrawImageOptions{GeoM: g.targetPos})

	firePos := g.targetPos
	firePos.Translate(
		float64(playerSprite.Bounds().Dx())/2-float64(fireSprite.Bounds().Dx())/2,
		float64(playerSprite.Bounds().Dy()),
	)
	screen.DrawImage(fireSprite, &ebiten.DrawImageOptions{GeoM: firePos})
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
