package system

import (
	"math/rand"

	"m3g4p0p/game/component"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/features/transform"
	"github.com/yohamta/donburi/filter"
)

type target struct {
	query *donburi.Query
}

func NewTarget() target {
	return target{donburi.NewQuery(filter.Contains(
		component.Target,
		transform.Transform,
	))}
}

func (t target) Update(ecs *ecs.ECS) {
	speed := 1 / float64(ebiten.TPS())

	for entry := range t.query.Iter(ecs.World) {
		if _, ok := transform.GetParent(entry); ok {
			continue
		}

		pos := transform.WorldPosition(entry)
		target := component.Target.Get(entry)
		delta := target.Sub(pos).MulScalar(speed)
		rot := delta.Angle(math.Vec2{}) - math.ToRadians(90)
		transform.SetWorldPosition(entry, pos.Add(delta))
		transform.SetWorldRotation(entry, rot)

		if fire, ok := transform.FindChildWithComponent(entry, component.Fire); ok {
			alpha := target.Distance(pos)*speed + rand.Float64()/10
			component.Fire.SetValue(fire, component.FireData(alpha))
		}
	}
}
