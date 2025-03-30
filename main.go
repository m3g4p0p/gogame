package main

import (
	"embed"
	"log"
	"math"

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
		g.targetPos = util.CursorPosition()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	cursorPos := util.CursorPosition()
	angle := cursorPos.Angle(g.targetPos) - math.Pi/2
	distance := cursorPos.Distance(g.targetPos)
	ebitenutil.DebugPrint(screen, g.targetPos.String())
	op := util.RotateCenter(playerSprite, angle, util.Vector{})
	op.GeoM.Translate(g.targetPos.X, g.targetPos.Y)
	screen.DrawImage(playerSprite, op)

	fireOp := util.RotateCenter(
		fireSprite,
		angle,
		util.Vector{
			Y: float64(playerSprite.Bounds().Dy()) * 0.7,
		},
	)

	fireOp.ColorScale.ScaleAlpha(float32(distance) / 100)
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
