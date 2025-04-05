package main

import (
	"embed"
	"errors"
	"log"
	"runtime"

	"m3g4p0p/game/component"
	"m3g4p0p/game/factory"
	"m3g4p0p/game/util"
	"m3g4p0p/game/vec2"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	vector "github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/features/transform"
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
	world      donburi.World
	playerPos  vec2.Vector
	targetPos  vec2.Vector
	targetVec2 vector.Vec2
}

func (g *Game) updateTransform() {
	query := donburi.NewQuery(filter.Contains(
		transform.Transform,
	))

	speed := 1 / float64(ebiten.TPS())

	for entry := range query.Iter(g.world) {
		if _, ok := transform.GetParent(entry); ok {
			continue
		}

		pos := transform.WorldPosition(entry)
		delta := g.targetVec2.Sub(pos).MulScalar(speed)
		transform.SetWorldPosition(entry, pos.Add(delta))
		transform.LookAt(entry, g.targetVec2)
		logger.Print(pos)
	}
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyC) && ebiten.IsKeyPressed(ebiten.KeyControl) {
		return errors.New("exit")
	}

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		g.targetVec2 = util.CursorPositionVec2()
	}

	g.updateTransform()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	query := donburi.NewQuery(filter.Contains(
		transform.Transform,
		component.Sprite,
	))

	for entry := range query.Iter(g.world) {
		pos := transform.WorldPosition(entry)
		rot := transform.WorldRotation(entry)
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(pos.X, pos.Y)
		op.GeoM.Rotate(rot)

		sprite := component.Sprite.Get(entry)
		screen.DrawImage(sprite, op)
	}

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
	player := factory.CreateShip(world, playerSprite, center.X, center.Y)
	factory.CreateFire(world, fireSprite, player)

	return &Game{world, center, center, vector.NewVec2(center.Values())}
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
