package types

import "github.com/hajimehoshi/ebiten/v2"

// Sprite represents an image with Bounds and a position
// that can be drawn by the Renderer
type Sprite interface {
	Bounds
	Image() Image
}

// UpdateableSprite represents a sprite that has an update function
type UpdateableSprite interface {
	Sprite
	Update(delta int)
}

//  ScaledSprite represents a sprite that has should be scaled when drawing
type ScaledSprite interface {
	Scale() (float64, float64)
}

type ColorScaleSprite interface {
	Sprite
	Color() ebiten.ColorScale
}

// ShaderSprite represents a sprite that should use a shader when drawing
type ShaderSprite interface {
	Sprite
	Shader() (Shader, map[string]any)
}
