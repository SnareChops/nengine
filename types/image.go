package types

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Image interface {
	Size() (w, h int)
	Clear()
	Fill(color.Color)
	Dx() int
	Dy() int
	DrawImage(Image, *DrawImageOptions)
	DrawRectShader(w, h int, shader Shader, options *DrawRectShaderOptions)
	SubImage(x, y, w, h int) Image
	Raw() *ebiten.Image
}

type DrawImageOptions = ebiten.DrawImageOptions
type DrawRectShaderOptions = ebiten.DrawRectShaderOptions
