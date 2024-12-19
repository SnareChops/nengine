package image

import (
	"image"
	"image/color"

	"github.com/SnareChops/nengine/types"
	"github.com/hajimehoshi/ebiten/v2"
)

type img ebiten.Image

var NewImage func(w, h int) types.Image = func(w, h int) types.Image {
	i := ebiten.NewImage(w, h)
	return (*img)(i)
}

var NewImageFromImage func(image.Image) types.Image = func(src image.Image) types.Image {
	return (*img)(ebiten.NewImageFromImage(src))
}

func Wrap(i *ebiten.Image) types.Image {
	return (*img)(i)
}

func (i *img) Size() (w, h int) {
	return (*ebiten.Image)(i).Size()
}

func (i *img) Clear() {
	(*ebiten.Image)(i).Clear()
}

func (i *img) Fill(clr color.Color) {
	(*ebiten.Image)(i).Fill(clr)
}

func (i *img) Dx() int {
	return (*ebiten.Image)(i).Bounds().Dx()
}

func (i *img) Dy() int {
	return (*ebiten.Image)(i).Bounds().Dy()
}

func (i *img) DrawImage(image types.Image, options *types.DrawImageOptions) {
	(*ebiten.Image)(i).DrawImage((*ebiten.Image)(image.(*img)), options)
}

func (i *img) DrawRectShader(w, h int, shader types.Shader, options *types.DrawRectShaderOptions) {
	(*ebiten.Image)(i).DrawRectShader(w, h, shader.(*ebiten.Shader), options)
}

func (i *img) SubImage(x, y, w, h int) types.Image {
	goImg := (*ebiten.Image)(i).SubImage(image.Rect(x, y, x+w, y+h))
	eImg := goImg.(*ebiten.Image)
	return (*img)(eImg)
}

func (i *img) Raw() *ebiten.Image {
	return (*ebiten.Image)(i)
}
