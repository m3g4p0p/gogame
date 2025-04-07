package component

import (
	"m3g4p0p/game/vec2"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"
)

type FireData float64

var (
	Position = donburi.NewComponentType[vec2.Vector]()
	Sprite   = donburi.NewComponentType[ebiten.Image]()
	Target   = donburi.NewComponentType[math.Vec2]()
	Fire     = donburi.NewComponentType[FireData]()
)
