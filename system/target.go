package system

import (
	"math/rand"

	"m3g4p0p/game/component"

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
	speed := ecs.Time.DeltaTime().Seconds()

	for entry := range t.query.Iter(ecs.World) {
		pos := transform.WorldPosition(entry)
		target := component.Target.Get(entry)
		delta := target.Sub(pos).MulScalar(speed)
		rot := target.Angle(pos) - math.ToRadians(90)
		transform.SetWorldPosition(entry, pos.Add(delta))
		transform.SetWorldRotation(entry, rot)

		if fire, ok := transform.FindChildWithComponent(entry, component.Fire); ok {
			alpha := target.Distance(pos)*speed + rand.Float64()/10
			component.Fire.SetValue(fire, component.FireData(alpha))
		}
	}
}
