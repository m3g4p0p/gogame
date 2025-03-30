package util

import "github.com/hajimehoshi/ebiten/v2"

func RotateCenter(sprite *ebiten.Image, theta float64) *ebiten.DrawImageOptions {
	bounds := sprite.Bounds()
	dx, dy := float64(bounds.Dx()), float64(bounds.Dy())
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(-dx/2, -dy/2)
	op.GeoM.Rotate(theta)
	op.GeoM.Translate(dx/2, dy/2)

	return op
}
