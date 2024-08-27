package rendering

import (
	"github.com/SnareChops/nengine/types"
	"github.com/hajimehoshi/ebiten/v2"
)

type DrawableSprite interface {
	Dx() int
	Dy() int
	DrawOptions(types.Camera) *ebiten.DrawImageOptions
	Image() *ebiten.Image
}

func DrawSprite(dest *ebiten.Image, sprite DrawableSprite, camera types.Camera) {
	image := sprite.Image()
	if image != nil {
		options := sprite.DrawOptions(camera)
		if scaled, ok := sprite.(types.ScaledSprite); ok {
			options.GeoM.Scale(scaled.Scale())
		}
		if shader, uniforms, ok := shouldUseShader(sprite); ok {
			op := &ebiten.DrawRectShaderOptions{}
			op.GeoM = options.GeoM
			op.Uniforms = uniforms
			op.Images = [4]*ebiten.Image{image}
			dest.DrawRectShader(sprite.Dx(), sprite.Dy(), shader, op)
		} else {
			dest.DrawImage(image, options)
		}
	}
}

func shouldUseShader(sprite DrawableSprite) (*ebiten.Shader, map[string]any, bool) {
	if s, ok := sprite.(types.ShaderSprite); ok {
		if shader, uniforms := s.Shader(); shader != nil {
			return shader, uniforms, true
		}
	}
	return nil, nil, false
}
