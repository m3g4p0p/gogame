package main

import (
	"embed"
	"errors"
	"log"
	"runtime"

	"m3g4p0p/game/component"
	"m3g4p0p/game/factory"
	"m3g4p0p/game/util"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
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

func (g *Game) updateTransform() {
	query := donburi.NewQuery(filter.Contains(
		transform.Transform,
		component.Target,
	))

	speed := 1 / float64(ebiten.TPS())

	for entry := range query.Iter(g.world) {
		if _, ok := transform.GetParent(entry); ok {
			continue
		}

		pos := transform.WorldPosition(entry)
		target := component.Target.Get(entry)
		delta := target.Sub(pos).MulScalar(speed)
		rot := delta.Angle(math.Vec2{}) - math.ToRadians(90)
		transform.SetWorldPosition(entry, pos.Add(delta))
		transform.SetWorldRotation(entry, rot)
		logger.Print(target)

		if fire, ok := transform.FindChildWithComponent(entry, component.Fire); ok {
			alpha := target.Distance(pos) * speed
			component.Fire.SetValue(fire, component.FireData(alpha))
		}
	}
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyC) && ebiten.IsKeyPressed(ebiten.KeyControl) {
		return errors.New("exit")
	}

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		g.updateTarget()
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
	width, height := ebiten.WindowSize()
	center := math.NewVec2(float64(width)/2, float64(height)/2)
	player := factory.CreateShip(world, playerSprite, center, component.Player)
	factory.CreateFire(world, fireSprite, player)

	return &Game{world}
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
