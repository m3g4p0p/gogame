package main

import (
	"embed"
	"errors"
	"fmt"
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
	playerPos util.Vector
	targetPos util.Vector
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyC) && ebiten.IsKeyPressed(ebiten.KeyControl) {
		return errors.New("exit")
	}

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		g.targetPos = util.CursorPosition()
	}

	speed := 1 / float64(ebiten.TPS())
	delta := g.targetPos.Sub(g.playerPos)
	g.playerPos = g.playerPos.Add(delta.Scale(speed))

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	angle := util.CursorPosition().Angle(g.playerPos) - math.Pi/2
	distance := g.targetPos.Distance(g.playerPos)
	ebitenutil.DebugPrint(screen, fmt.Sprint(g.targetPos))

	op := util.RotateCenter(playerSprite, angle, util.Vector{})
	op.GeoM.Translate(g.playerPos.X, g.playerPos.Y)
	screen.DrawImage(playerSprite, op)

	fireOp := util.RotateCenter(
		fireSprite,
		angle,
		util.Vector{
			Y: float64(playerSprite.Bounds().Dy()) * 0.7,
		},
	)

	fireOp.ColorScale.ScaleAlpha(float32(distance) / 100)
	fireOp.GeoM.Translate(g.playerPos.X, g.playerPos.Y)
	screen.DrawImage(fireSprite, fireOp)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	ebiten.SetFullscreen(true)

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
