package util

import (
	"m3g4p0p/game/vec2"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi/features/math"
)

func RotateCenter(img *ebiten.Image, theta float64, offset vec2.Vector) *ebiten.DrawImageOptions {
	bounds := img.Bounds()
	dx, dy := float64(bounds.Dx()), float64(bounds.Dy())
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
