package component

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"
)

type FireData float64

var (
	Sprite = donburi.NewComponentType[ebiten.Image]()
	Target = donburi.NewComponentType[math.Vec2]()
	Fire   = donburi.NewComponentType[FireData]()
	Player = donburi.NewTag("Player")
)
