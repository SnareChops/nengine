package nengine

import (
	"image/color"

	"github.com/SnareChops/nengine/types"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Line struct {
	x1, y1, x2, y2 float32
}

func (self Line) Points() (Position, Position) {
	return Point(float64(self.x1), float64(self.y1)), Point(float64(self.x2), float64(self.y2))
}

func NewLine[T ~int | ~float64](x1, y1, x2, y2 T) Line {
	return Line{
		x1: float32(x1),
		y1: float32(y1),
		x2: float32(x2),
		y2: float32(y2),
	}
}

func StrokeLine(dest types.Image, line Line, strokeWidth float32, color color.Color, antialias bool, camera types.Camera) {
	StrokeLineRaw(dest, line.x1, line.y1, line.x2, line.y2, strokeWidth, color, antialias, camera)
}

func StrokeLineBetween(dest types.Image, start, end types.Position, strokeWidth float32, color color.Color, antialias bool, camera types.Camera) {
	StrokeLineRaw(dest, start.X(), start.Y(), end.X(), end.Y(), strokeWidth, color, antialias, camera)
}

func StrokeLineRaw[T ~int | ~float32 | ~float64](dest types.Image, x1, y1, x2, y2 T, strokeWidth float32, color color.Color, antialias bool, camera types.Camera) {
	if camera != nil {
		sx, sy := camera.WorldToScreenPos(float64(x1), float64(y1))
		ex, ey := camera.WorldToScreenPos(float64(x2), float64(y2))
		vector.StrokeLine(dest.Raw(), float32(sx), float32(sy), float32(ex), float32(ey), strokeWidth, color, antialias)
		return
	}
	vector.StrokeLine(dest.Raw(), float32(x1), float32(y1), float32(x2), float32(y2), strokeWidth, color, antialias)
}
