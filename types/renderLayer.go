package types

type RenderLayer interface {
	Order() int
	Draw(screen Image)
}

type SpriteRenderLayer interface {
	RenderLayer
	Sprites() []Sprite
	AddSprite(sprite Sprite)
	RemoveSprite(sprite Sprite)
}
