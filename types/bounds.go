package types

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Box interface {
	Position
	Size() (int, int)
	SetSize(w, h int)
	Resize(w, h int)
	Offset() (float64, float64)
	SetOffset(x, y float64)
	SetAnchor(h, v int)
	Rotation() float64
	SetRotation(float64)
	Flip(h, v bool)
	Dx() int
	Dy() int
	Min() (float64, float64)
	Mid() (float64, float64)
	Max() (float64, float64)
	MidX() float64
	MidY() float64
	MaxX() float64
	MaxY() float64
}

// Bounds represents a rectangle
// Includes useful utilities for working with the bounds
type Bounds interface {
	Box
	DrawOptions(sx, sy float64, camera Camera) *ebiten.DrawImageOptions
}

type PhysicsBounds interface {
	Bounds
	Velocity() (float64, float64)
	SetRawVelocity(x, y float64)
	SetVelocity(angle, magnitude float64)
	Acceleration() (float64, float64)
	SetAcceleration(angle, magnitude float64)
	Update(delta int)
}
