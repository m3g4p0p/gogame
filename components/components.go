package components

import (
	"m3g4p0p/game/vec2"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
)

var (
	Position = donburi.NewComponentType[vec2.Vector]()
	Sprite   = donburi.NewComponentType[ebiten.Image]()
)
