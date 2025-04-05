package util

import (
	"m3g4p0p/game/vec2"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi/features/math"
)

func CenterOffset(img *ebiten.Image) math.Vec2 {
	bounds := img.Bounds()
	dx := float64(bounds.Dx())
	dy := float64(bounds.Dy())
	return math.NewVec2(-dx/2, -dy/2)
}

func RotateCenter(img *ebiten.Image, theta float64, offset vec2.Vector) *ebiten.DrawImageOptions {
	dx, dy := CenterOffset(img).XY()
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(-dx/2+offset.X, -dy/2+offset.Y)
	op.GeoM.Rotate(theta)

	return op
}

func CursorPosition() vec2.Vector {
	x, y := ebiten.CursorPosition()
	return vec2.Vector{X: float64(x), Y: float64(y)}
}

func CursorPositionVec2() math.Vec2 {
	x, y := ebiten.CursorPosition()
	return math.NewVec2(float64(x), float64(y))
}
