package system

import (
	"m3g4p0p/game/component"
	"m3g4p0p/game/util"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/features/transform"
	"github.com/yohamta/donburi/filter"
)

type render struct {
	query *donburi.Query
}

func NewRender() render {
	return render{donburi.NewQuery(filter.Contains(
		component.Sprite,
		transform.Transform,
	))}
}

func (r render) Draw(ecs *ecs.ECS, screen *ebiten.Image) {
	for entry := range r.query.Iter(ecs.World) {
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
}
