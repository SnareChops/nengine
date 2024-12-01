package nengine

import (
	"github.com/SnareChops/nengine/bounds"
	"github.com/SnareChops/nengine/types"
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
	image types.Image
}

// Init sets the initial state of the SimpleStruct
func (self *SimpleSprite) Init(image types.Image) *SimpleSprite {
	self.image = image
	self.RawBounds = new(bounds.Raw).Init(self.image.Size())
	return self
}

// Image returns the image for drawing
func (self *SimpleSprite) Image() types.Image {
	return self.image
}

type SimpleShaderSprite struct {
	*SimpleSprite
	shader   types.Shader
	uniforms map[string]any
}

func (self *SimpleShaderSprite) Init(image types.Image, shader types.Shader, uniforms map[string]any) *SimpleShaderSprite {
	self.SimpleSprite = new(SimpleSprite).Init(image)
	self.shader = shader
	self.uniforms = uniforms
	return self
}

func (self *SimpleShaderSprite) Shader() (types.Shader, map[string]any) {
	return self.shader, self.uniforms
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

func (self *SourceSprite) Image() types.Image {
	return self.SimpleSprite.Image()
}
