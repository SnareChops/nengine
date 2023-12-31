package types

import "github.com/hajimehoshi/ebiten/v2"

type RenderLayer interface {
	Order() int
	Draw(screen *ebiten.Image)
}

type SpriteRenderLayer interface {
	RenderLayer
	Sprites() []Sprite
	AddSprite(sprite Sprite)
	RemoveSprite(sprite Sprite)
}
