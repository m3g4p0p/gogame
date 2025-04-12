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
	ecslib "github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/features/math"
)

//go:embed assets/*
var assets embed.FS
var logger = util.NewLogger(log.Lshortfile)

var (
	playerSprite = util.Must(util.LoadImage(assets, "assets/playerShip1_blue.png"))
	fireSprite   = util.Must(util.LoadImage(assets, "assets/Effects/fire11.png"))
	meteorSprite = util.Must(util.LoadImage(assets, "assets/Meteors/meteorGrey_small1.png"))
)

type Game struct {
	ecs *ecslib.ECS
}

func (g *Game) updateTarget() {
	player := component.Player.MustFirst(g.ecs.World)
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
	g.ecs.Draw(screen)
	logger.Flush(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func newGame() *Game {
	world := donburi.NewWorld()
	ecs := ecslib.NewECS(world)
	ecs.AddSystem(system.NewTarget().Update)
	ecs.AddRenderer(ecslib.LayerDefault, system.NewRender().Draw)

	width, height := ebiten.WindowSize()
	center := math.NewVec2(float64(width)/2, float64(height)/2)
	player := factory.CreateTransform(world, playerSprite, center, component.Player)
	offset := math.NewVec2(0, float64(playerSprite.Bounds().Dy())*0.7)
	factory.CreateTransformChild(world, fireSprite, player, offset, component.Fire)

	return &Game{ecs}
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
