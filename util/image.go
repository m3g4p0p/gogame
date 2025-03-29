package util

import (
	"embed"
	"image"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

func LoadImage(fs embed.FS, name string) (*ebiten.Image, error) {
	f, err := fs.Open(name)

	if err != nil {
		return nil, err
	}

	defer f.Close()

	if img, _, err := image.Decode(f); err != nil {
		return nil, err
	} else {
		return ebiten.NewImageFromImage(img), nil
	}
}
