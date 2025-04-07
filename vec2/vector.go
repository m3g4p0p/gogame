package vec2

import (
	"fmt"
	"math"
)

type Vector struct {
	X, Y float64
}

func (v Vector) Add(t Vector) Vector {
	return Vector{v.X + t.X, v.Y + t.Y}
}

func (v Vector) Sub(t Vector) Vector {
	return Vector{v.X - t.X, v.Y - t.Y}
}

func (v Vector) Scale(t float64) Vector {
	return Vector{v.X * t, v.Y * t}
}

func (v Vector) Angle(t Vector) float64 {
	return math.Atan2(t.Y-v.Y, t.X-v.X)
}

func (v Vector) Distance(t Vector) float64 {
	return math.Hypot(t.X-v.X, t.Y-v.Y)
}

func (v Vector) String() string {
	return fmt.Sprintf("(%v, %v)", v.X, v.Y)
}

func (v Vector) Values() (float64, float64) {
	return v.X, v.Y
}
