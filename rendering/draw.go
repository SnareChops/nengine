package rendering

import (
	"github.com/SnareChops/nengine/types"
)

func DrawAt[T ~int | float32 | float64](dest types.Image, src types.Image, x, y T) {
	options := &types.DrawImageOptions{}
	options.GeoM.Translate(float64(x), float64(y))
	dest.DrawImage(src, options)
}

func GridDraw(image types.Image, cell types.Image, index int) {
	width := image.Dx() / cell.Dx()
	x := index % width
	y := index / width
	options := &types.DrawImageOptions{}
	options.GeoM.Translate(float64(x*cell.Dx()), float64(y*cell.Dy()))
	image.DrawImage(cell, options)
}
