package nengine

import (
	"github.com/SnareChops/nengine/bounds"
	"github.com/hajimehoshi/ebiten/v2"
)

// SimpleSprite is a convenience struct for
// "just drawing a raw image to the screen"
// Note: Unless your intention is to just draw a
// static image to the screen with no behavior
// then it is **strongly** recommended to implement
// the Sprite interface in your own struct instead
// of using this
type SimpleSprite struct {
	*RawBounds
	image *ebiten.Image
}

// Init sets the initial state of the SimpleStruct
func (self *SimpleSprite) Init(image *ebiten.Image) *SimpleSprite {
	self.image = image
	self.RawBounds = new(bounds.Raw).Init(self.image.Size())
	return self
}

// Image returns the image for drawing
func (self *SimpleSprite) Image() *ebiten.Image {
	return self.image
}

type SourceSprite struct {
	*SimpleSprite
	SpriteSource
}

func (self *SourceSprite) Init(source SpriteSource) *SourceSprite {
	self.SpriteSource = source
	self.SimpleSprite = new(SimpleSprite).Init(source.Image())
	return self
}

func (self *SourceSprite) Reload() {
	self.image = self.SpriteSource.Image()
}

func (self *SourceSprite) Image() *ebiten.Image {
	return self.SimpleSprite.Image()
}

func DrawSpriteWithShader(dest *ebiten.Image, sprite Sprite, shader *ebiten.Shader, camera Camera, uniforms map[string]interface{}) {
	o := sprite.DrawOptions(camera)
	options := &ebiten.DrawRectShaderOptions{
		GeoM:       o.GeoM,
		ColorScale: o.ColorScale,
		Uniforms:   uniforms,
		Images:     [4]*ebiten.Image{sprite.Image()},
	}
	dest.DrawRectShader(sprite.Dx(), sprite.Dy(), shader, options)
}
