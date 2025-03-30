package main

import (
	"embed"
	"fmt"
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

type Vector struct {
	X, Y float64
}

func (v Vector) String() string {
	return fmt.Sprintf("(%v, %v)", v.X, v.Y)
}

type Game struct {
	targetPos Vector
	angle     float64
}

func (g *Game) Update() error {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		g.targetPos = Vector{float64(x), float64(y)}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, g.targetPos.String())
	op := util.RotateCenter(playerSprite, 0.1)
	op.GeoM.Translate(g.targetPos.X, g.targetPos.Y)
	screen.DrawImage(playerSprite, op)
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
