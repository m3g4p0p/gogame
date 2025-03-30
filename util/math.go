package util

import "github.com/hajimehoshi/ebiten/v2"

func RotateCenter(img *ebiten.Image, theta, offsetX, offsetY float64) *ebiten.DrawImageOptions {
	bounds := img.Bounds()
	dx, dy := float64(bounds.Dx()), float64(bounds.Dy())
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(-dx/2+offsetX, -dy/2+offsetY)
	op.GeoM.Rotate(theta)
	op.GeoM.Translate(dx/2, dy/2)

	return op
}
