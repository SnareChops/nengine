package types

import "github.com/hajimehoshi/ebiten/v2"

// Sprite represents an image with Bounds and a position
// that can be drawn by the Renderer
type Sprite interface {
	Bounds
	Image() *ebiten.Image
}

// UpdateableSprite represents a sprite that has an update function
type UpdateableSprite interface {
	Sprite
	Update(delta int)
}

type ColorScaleSprite interface {
	Sprite
	Color() ebiten.ColorScale
}
