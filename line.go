package nengine

import (
	"image/color"

	"github.com/SnareChops/nengine/types"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Line struct {
	x1, y1, x2, y2 float32
}

func NewLine[T ~int | ~float64](x1, y1, x2, y2 T) Line {
	return Line{
		x1: float32(x1),
		y1: float32(y1),
		x2: float32(x2),
		y2: float32(y2),
	}
}

func DrawLine(dest *ebiten.Image, line Line, size float32, color color.Color, antialias bool, camera types.Camera) {
	if camera != nil {
		x1, y1 := camera.WorldToScreenPos(float64(line.x1), float64(line.y1))
		x2, y2 := camera.WorldToScreenPos(float64(line.x2), float64(line.y2))
		vector.StrokeLine(dest, float32(x1), float32(y1), float32(x2), float32(y2), size, color, antialias)
	} else {
		vector.StrokeLine(dest, line.x1, line.y1, line.x2, line.y2, size, color, antialias)
	}
}
