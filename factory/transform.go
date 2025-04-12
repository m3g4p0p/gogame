package factory

import (
	"m3g4p0p/game/component"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/features/transform"
)

func CreateTransform(
	world donburi.World,
	sprite *ebiten.Image,
	position math.Vec2,
	components ...donburi.IComponentType,
) *donburi.Entry {
	components = append(components, transform.Transform, component.Sprite)
	entry := world.Entry(world.Create(components...))
	transform.Transform.SetValue(entry, transform.TransformData{
		LocalPosition: position,
	})
	component.Sprite.Set(entry, sprite)
	return entry
}

func CreateTransformChild(
	world donburi.World,
	sprite *ebiten.Image,
	parent *donburi.Entry,
	offset math.Vec2,
	components ...donburi.IComponentType,
) *donburi.Entry {
	entry := CreateTransform(world, sprite, offset, components...)
	transform.AppendChild(parent, entry, false)
	return entry
}
