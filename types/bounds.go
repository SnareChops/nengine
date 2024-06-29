package types

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// Bounds represents a rectangle
// Includes useful utilities for working with the bounds
type Bounds interface {
	Position
	Collidable
	PosOf(h, v int) (x, y float64)
	RawPos() (x, y float64)
	Anchor() (h, v int)
	SetAnchor(h, v int)
	Offset() (x, y float64)
	SetOffset(x, y float64)
	Size() (w, h int)
	SetSize(w, h int)
	Dx() int
	Dy() int
	Rotation() float64
	SetRotation(theta float64)
	Scale() (x, y float64)
	SetScale(x, y float64)
	ScaleTo(w, h int)
	Mid() (x, y float64)
	MinX() float64
	MinY() float64
	MaxX() float64
	MaxY() float64
	DrawOptions(camera Camera) *ebiten.DrawImageOptions
}

type Collidable interface {
	X() float64
	Y() float64
	Min() (x, y float64)
	Max() (x, y float64)
	IsWithin(x, y float64) bool
	DoesCollide(Collidable) bool
}
