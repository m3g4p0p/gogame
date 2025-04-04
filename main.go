package main

import (
	"embed"
	"errors"
	"log"
	"math"
	"runtime"

	"m3g4p0p/game/components"
	"m3g4p0p/game/util"
	"m3g4p0p/game/vec2"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
)

//go:embed assets/*
var assets embed.FS
var logger = util.NewLogger(log.Lshortfile)

var (
	playerSprite = util.Must(util.LoadImage(assets, "assets/playerShip1_blue.png"))
	fireSprite   = util.Must(util.LoadImage(assets, "assets/Effects/fire09.png"))
)

type Game struct {
	world     donburi.World
	playerPos vec2.Vector
	targetPos vec2.Vector
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

	query := donburi.NewQuery(filter.Contains(
		components.Position,
	))

	for entry := range query.Iter(g.world) {
		pos := components.Position.Get(entry)
		delta := g.targetPos.Sub(*pos)
		newPos := pos.Add(delta.Scale(speed))
		components.Position.Set(entry, &newPos)
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	angle := util.CursorPosition().Angle(g.playerPos) - math.Pi/2
	distance := g.targetPos.Distance(g.playerPos)

	op := util.RotateCenter(playerSprite, angle, vec2.Vector{})
	op.GeoM.Translate(g.playerPos.X, g.playerPos.Y)
	screen.DrawImage(playerSprite, op)

	fireOp := util.RotateCenter(
		fireSprite,
		angle,
		vec2.Vector{
			Y: float64(playerSprite.Bounds().Dy()) * 0.7,
		},
	)

	fireOp.ColorScale.ScaleAlpha(float32(distance) / 100)
	fireOp.GeoM.Translate(g.playerPos.X, g.playerPos.Y)
	screen.DrawImage(fireSprite, fireOp)

	logger.Flush(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func newGame() *Game {
	width, height := ebiten.WindowSize()

	center := vec2.Vector{
		X: float64(width) / 2,
		Y: float64(height) / 2,
	}

	world := donburi.NewWorld()
	player := world.Create(components.Position)
	entry := world.Entry(player)
	donburi.Add(entry, components.Sprite, playerSprite)

	return &Game{world, center, center}
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")

	if runtime.GOOS == "js" {
		ebiten.SetFullscreen(true)
	}

	if err := ebiten.RunGame(newGame()); err != nil {
		log.Fatal(err)
	}
}
