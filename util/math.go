package util

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

type Vector struct {
	X, Y float64
}

func (v Vector) String() string {
	return fmt.Sprintf("(%v, %v)", v.X, v.Y)
}

func RotateCenter(img *ebiten.Image, theta, offsetX, offsetY float64) *ebiten.DrawImageOptions {
	bounds := img.Bounds()
	dx, dy := float64(bounds.Dx()), float64(bounds.Dy())
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(-dx/2+offsetX, -dy/2+offsetY)
	op.GeoM.Rotate(theta)
	op.GeoM.Translate(dx/2, dy/2)

	return op
}
