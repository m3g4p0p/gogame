package factory

import (
	"m3g4p0p/game/component"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/features/transform"
)

func CreateShip(world donburi.World, sprite *ebiten.Image, x, y float64) *donburi.Entry {
	ship := world.Entry(world.Create(transform.Transform, component.Sprite))
	transform.Transform.SetValue(ship, transform.TransformData{
		LocalPosition: math.Vec2{X: x, Y: y},
	})
	component.Sprite.Set(ship, sprite)
	return ship
}

func CreateFire(world donburi.World, sprite *ebiten.Image, ship *donburi.Entry) *donburi.Entry {
	fire := world.Entry(world.Create(transform.Transform, component.Sprite))
	transform.Transform.SetValue(fire, transform.TransformData{
		LocalPosition: math.Vec2{X: 0, Y: float64(sprite.Bounds().Dy())},
	})
	component.Sprite.Set(fire, sprite)
	transform.SetParent(fire, ship, false)
	return fire
}
