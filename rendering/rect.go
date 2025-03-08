package rendering

import (
	"image/color"

	"github.com/SnareChops/nengine/types"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func StrokeRect(dest types.Image, corner1, corner2 types.Position, strokeWidth float32, color color.Color, camera types.Camera) {
	StrokeRectRaw(dest, corner1.X(), corner1.Y(), corner2.X(), corner2.Y(), strokeWidth, color, camera)
}

func StrokeBox(dest types.Image, box types.Box, strokeWidth float32, color color.Color, camera types.Camera) {
	StrokeRectRaw(dest, box.MinX(), box.MinY(), box.MaxX(), box.MaxY(), strokeWidth, color, camera)
}

func StrokeRectRaw[T int | float64](dest types.Image, x1, y1, x2, y2 T, strokeWidth float32, color color.Color, camera types.Camera) {
	mx := min(x1, x2)
	my := min(y1, y2)
	Mx := max(x1, x2)
	My := max(y1, y2)
	if camera != nil {
		mx, my := camera.WorldToScreenPos(float64(mx), float64(my))
		Mx, My := camera.WorldToScreenPos(float64(Mx), float64(My))
		vector.StrokeRect(dest.Raw(), float32(mx), float32(my), float32(Mx-mx), float32(My-my), strokeWidth, color, false)
		return
	}
	vector.StrokeRect(dest.Raw(), float32(mx), float32(my), float32(Mx-mx), float32(My-my), strokeWidth, color, true)
}

func StrokeLine(dest types.Image, start, end types.Position, strokeWidth float32, color color.Color, camera types.Camera) {
	mx := min(start.X(), end.X())
	my := min(start.Y(), end.Y())
	Mx := max(start.X(), end.X())
	My := max(start.Y(), end.Y())
	if camera != nil {
		mx, my := camera.WorldToScreenPos(mx, my)
		Mx, My := camera.WorldToScreenPos(Mx, My)
		vector.StrokeLine(dest.Raw(), float32(mx), float32(my), float32(Mx), float32(My), strokeWidth, color, true)
		return
	}
	vector.StrokeLine(dest.Raw(), float32(mx), float32(my), float32(Mx), float32(My), strokeWidth, color, true)
}
