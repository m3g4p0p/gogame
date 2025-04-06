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
		LocalPosition: math.NewVec2(x, y),
	})
	component.Sprite.Set(ship, sprite)
	return ship
}

func CreateFire(world donburi.World, sprite *ebiten.Image, ship *donburi.Entry) *donburi.Entry {
	fire := world.Entry(world.Create(transform.Transform, component.Sprite))
	transform.Transform.SetValue(fire, transform.TransformData{
		LocalPosition: math.NewVec2(0, float64(sprite.Bounds().Dy())),
	})
	component.Sprite.Set(fire, sprite)
	transform.AppendChild(ship, fire, false)
	return fire
}
