package main

import (
	"embed"
	"errors"
	"log"
	"runtime"

	"m3g4p0p/game/component"
	"m3g4p0p/game/factory"
	"m3g4p0p/game/system"
	"m3g4p0p/game/util"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/features/math"
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
	world donburi.World
	ecs   *ecs.ECS
}

func (g *Game) updateTarget() {
	player := component.Player.MustFirst(g.world)
	target := util.CursorPositionVec2()

	if player.HasComponent(component.Target) {
		component.Target.SetValue(player, target)
	} else {
		donburi.Add(player, component.Target, &target)
	}
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyC) && ebiten.IsKeyPressed(ebiten.KeyControl) {
		return errors.New("exit")
	}

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		g.updateTarget()
	}

	g.ecs.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	query := donburi.NewQuery(filter.Contains(
		transform.Transform,
		component.Sprite,
	))

	for entry := range query.Iter(g.world) {
		op := &ebiten.DrawImageOptions{}
		pos := util.TranslatedWorldPosition(entry)
		rot := transform.WorldRotation(entry)
		sprite := component.Sprite.Get(entry)
		op.GeoM.Translate(util.CenterOffset(sprite).XY())
		op.GeoM.Rotate(rot)
		op.GeoM.Translate(pos.XY())

		if entry.HasComponent(component.Fire) {
			alpha := component.Fire.GetValue(entry)
			op.ColorScale.ScaleAlpha(float32(alpha))
		}

		screen.DrawImage(sprite, op)
	}

	logger.Flush(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func newGame() *Game {
	world := donburi.NewWorld()
	ecs := ecs.NewECS(world)
	ecs.AddSystem(system.NewTarget().Update)
	width, height := ebiten.WindowSize()
	center := math.NewVec2(float64(width)/2, float64(height)/2)
	player := factory.CreateShip(world, playerSprite, center, component.Player)
	factory.CreateFire(world, fireSprite, player)

	return &Game{world, ecs}
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
