package image

import (
	"image"
	"image/color"

	"github.com/SnareChops/nengine/types"
	"github.com/hajimehoshi/ebiten/v2"
)

func EnableMocks() {
	NewImage = mockNewImage
	NewImageFromImage = mockNewImageFromImage
}

type imageMock struct {
	w, h int
}

func mockNewImage(w, h int) types.Image {
	return &imageMock{w: w, h: h}
}

func mockNewImageFromImage(src image.Image) types.Image {
	return &imageMock{w: src.Bounds().Dx(), h: src.Bounds().Dy()}
}

func (m imageMock) Dx() int {
	return m.w
}

func (m imageMock) Dy() int {
	return m.h
}

func (m imageMock) Size() (w, h int) {
	return m.w, m.h
}

func (m imageMock) Fill(clr color.Color) {

}

func (m imageMock) Clear() {

}

func (m imageMock) SubImage(x, y, w, h int) types.Image {
	return &imageMock{}
}

func (m imageMock) DrawImage(src types.Image, options *types.DrawImageOptions) {

}

func (m imageMock) DrawRectShader(w, h int, shader types.Shader, options *types.DrawRectShaderOptions) {

}

func (m imageMock) Raw() *ebiten.Image {
	return &ebiten.Image{}
}
