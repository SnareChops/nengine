package nengine

import (
	"image/color"

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

func DrawLine(dest *ebiten.Image, line Line, size float32, color color.Color, antialias bool) {
	vector.StrokeLine(dest, line.x1, line.y1, line.x2, line.y2, size, color, antialias)
}
