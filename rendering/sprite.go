package rendering

import (
	"github.com/SnareChops/nengine/types"
	"github.com/hajimehoshi/ebiten/v2"
)

type DrawableSprite interface {
	Dx() int
	Dy() int
	DrawOptions(sx, sy float64, camera types.Camera) *types.DrawImageOptions
	Image() types.Image
}

func DrawSprite(dest types.Image, sprite DrawableSprite, camera types.Camera) {
	image := sprite.Image()
	if image != nil {
		var sx, sy float64 = 1, 1
		if scaled, ok := sprite.(types.ScaledSprite); ok {
			sx, sy = scaled.Scale()
		}
		options := sprite.DrawOptions(sx, sy, camera)
		if shader, uniforms, ok := shouldUseShader(sprite); ok {
			op := &ebiten.DrawRectShaderOptions{}
			op.GeoM = options.GeoM
			op.Uniforms = uniforms
			op.Images = [4]*ebiten.Image{image.Raw()}
			dest.DrawRectShader(sprite.Dx(), sprite.Dy(), shader, op)
		} else {
			dest.DrawImage(image, options)
		}
	}
}

func DrawSpriteWithShader(dest types.Image, sprite DrawableSprite, shader types.Shader, uniforms map[string]any, camera types.Camera) {
	image := sprite.Image()
	if image != nil {
		var sx, sy float64 = 1, 1
		if scaled, ok := sprite.(types.ScaledSprite); ok {
			sx, sy = scaled.Scale()
		}
		op := sprite.DrawOptions(sx, sy, camera)
		options := &ebiten.DrawRectShaderOptions{}
		options.GeoM = op.GeoM
		options.Uniforms = uniforms
		options.Images = [4]*ebiten.Image{image.Raw()}
		dest.DrawRectShader(sprite.Dx(), sprite.Dy(), shader, options)
	}
}

func shouldUseShader(sprite DrawableSprite) (types.Shader, map[string]any, bool) {
	if s, ok := sprite.(types.ShaderSprite); ok {
		if shader, uniforms := s.Shader(); shader != nil {
			return shader, uniforms, true
		}
	}
	return nil, nil, false
}
