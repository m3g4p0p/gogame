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

var (
	playerSprite = util.Must(util.LoadImage(assets, "assets/playerShip1_blue.png"))
	fireSprite   = util.Must(util.LoadImage(assets, "assets/Effects/fire09.png"))
)

type Game struct {
	targetPos util.Vector
}

func (g *Game) Update() error {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		g.targetPos = util.Vector{X: float64(x), Y: float64(y)}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	angle := 0.1
	ebitenutil.DebugPrint(screen, g.targetPos.String())
	op := util.RotateCenter(playerSprite, angle, util.Vector{})
	op.GeoM.Translate(g.targetPos.X, g.targetPos.Y)
	screen.DrawImage(playerSprite, op)

	fireOp := util.RotateCenter(
		fireSprite,
		angle,
		util.Vector{
			X: float64(playerSprite.Bounds().Dx())/2 - float64(fireSprite.Bounds().Dx())/2,
			Y: float64(playerSprite.Bounds().Dy()),
		},
	)

	fireOp.GeoM.Translate(g.targetPos.X, g.targetPos.Y)
	screen.DrawImage(fireSprite, fireOp)
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
