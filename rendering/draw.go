package rendering

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func DrawAt[T ~int | float32 | float64](dest *ebiten.Image, src *ebiten.Image, x, y T) {
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(float64(x), float64(y))
	dest.DrawImage(src, options)
}

func GridDraw(image *ebiten.Image, cell *ebiten.Image, index int) {
	width := image.Bounds().Dx() / cell.Bounds().Dx()
	x := index % width
	y := index / width
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(float64(x*cell.Bounds().Dx()), float64(y*cell.Bounds().Dy()))
	image.DrawImage(cell, options)
}
