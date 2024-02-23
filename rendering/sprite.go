package rendering

import (
	"github.com/SnareChops/nengine/types"
	"github.com/hajimehoshi/ebiten/v2"
)

func DrawSprite(dest *ebiten.Image, sprite types.Sprite, camera types.Camera) {
	image := sprite.Image()
	if image != nil {
		if shader, uniforms, ok := shouldUseShader(sprite); ok {
			options := &ebiten.DrawRectShaderOptions{}
			options.GeoM = sprite.DrawOptions(camera).GeoM
			options.Uniforms = uniforms
			options.Images = [4]*ebiten.Image{image}
			dest.DrawRectShader(sprite.Dx(), sprite.Dy(), shader, options)
		} else {
			dest.DrawImage(image, sprite.DrawOptions(camera))
		}
	}
}

func shouldUseShader(sprite types.Sprite) (*ebiten.Shader, map[string]any, bool) {
	if s, ok := sprite.(types.ShaderSprite); ok {
		if shader, uniforms := s.Shader(); shader != nil {
			return shader, uniforms, true
		}
	}
	return nil, nil, false
}
