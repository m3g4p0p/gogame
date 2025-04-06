package util

import (
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/features/transform"
)

func TranslatedWorldPosition(entry *donburi.Entry) math.Vec2 {
	if parent, ok := transform.GetParent(entry); ok {
		offset := transform.GetTransform(entry).LocalPosition.Rotate(
			transform.WorldRotation(parent),
		)

		return TranslatedWorldPosition(parent).Add(offset)
	} else {
		return transform.WorldPosition(entry)
	}
}
