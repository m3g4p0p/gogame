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
var playerSprite = util.Must(util.LoadImage(assets, "assets/playerShip1_blue.png"))

type Player struct {
	pos *ebiten.GeoM
}

type Game struct {
	player    *Player
	targetPos *ebiten.GeoM
}

func (g *Game) Update() error {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		g.targetPos.Reset()
		g.targetPos.Translate(float64(x), float64(y))
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, fmt.Sprintf("%+v", g.player.pos))
	op := &ebiten.DrawImageOptions{GeoM: *g.targetPos}
	screen.DrawImage(playerSprite, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	game := &Game{player: &Player{&ebiten.GeoM{}}, targetPos: &ebiten.GeoM{}}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
