package util

import (
	"fmt"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Vector struct {
	X, Y float64
}

func (v Vector) Angle(t Vector) float64 {
	return math.Atan2(t.Y-v.Y, t.X-v.X)
}

func (v Vector) Distance(t Vector) float64 {
	return math.Sqrt(math.Pow(t.X-v.X, 2) + math.Pow(t.Y-v.Y, 2))
}

func (v Vector) String() string {
	return fmt.Sprintf("(%v, %v)", v.X, v.Y)
}

func RotateCenter(img *ebiten.Image, theta float64, offset Vector) *ebiten.DrawImageOptions {
	bounds := img.Bounds()
	dx, dy := float64(bounds.Dx()), float64(bounds.Dy())
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(-dx/2+offset.X, -dy/2+offset.Y)
	op.GeoM.Rotate(theta)

	return op
}

func CursorPosition() Vector {
	x, y := ebiten.CursorPosition()
	return Vector{float64(x), float64(y)}
}
