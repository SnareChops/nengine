package types

import (
	"math"
)

type Vector struct {
	X, Y float64
}

func NewVector(angle, magnitude float64) Vector {
	return Vector{math.Cos(angle) * magnitude, math.Sin(angle) * magnitude}
}

func (v Vector) Add(b Vector) Vector {
	return Vector{v.X + b.X, v.Y + b.Y}
}

func (v Vector) Sub(b Vector) Vector {
	return Vector{v.X - b.X, v.Y - b.Y}
}

func (v Vector) Normalize() Vector {
	magnitude := math.Sqrt(v.X*v.X + v.Y*v.Y)
	if magnitude == 0 {
		return v
	}
	return Vector{v.X / magnitude, v.Y / magnitude}
}

func (v Vector) Scale(scale float64) Vector {
	return Vector{v.X * scale, v.Y * scale}
}

func (v Vector) Distance(b Vector) float64 {
	diff := v.Sub(b)
	return math.Sqrt(diff.X*diff.X + diff.Y*diff.Y)
}
