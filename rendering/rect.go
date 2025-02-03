package rendering

import (
	"image/color"

	"github.com/SnareChops/nengine/types"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func StrokeRect(dest types.Image, corner1, corner2 types.Position, width float32, color color.Color) {
	minX := min(corner1.X(), corner2.X())
	minY := min(corner1.Y(), corner2.Y())
	maxX := max(corner1.X(), corner2.X())
	maxY := max(corner1.Y(), corner2.Y())
	vector.StrokeRect(dest.Raw(), float32(minX), float32(minY), float32(maxX-minX), float32(maxY-minY), width, color, false)
}

func StrokeBox(dest types.Image, box types.Box, strokeWidth float32, color color.Color, camera types.Camera) {
	x, y := box.Min()
	sx, sy := camera.WorldToScreenPos(x, y)

	vector.StrokeRect(dest.Raw(), float32(sx), float32(sy), float32(box.Dx()), float32(box.Dy()), strokeWidth, color, false)
}

func StrokeLine(dest types.Image, start, end types.Position, strokeWidth float32, color color.Color, camera types.Camera) {
	minX := min(start.X(), end.X())
	minY := min(start.Y(), end.Y())
	maxX := max(start.X(), end.X())
	maxY := max(start.Y(), end.Y())
	if camera != nil {
		minX, minY := camera.WorldToScreenPos(minX, minY)
		maxX, maxY := camera.WorldToScreenPos(maxX, maxY)
		vector.StrokeLine(dest.Raw(), float32(minX), float32(minY), float32(maxX), float32(maxY), strokeWidth, color, true)
		return
	}

	vector.StrokeLine(dest.Raw(), float32(minX), float32(minY), float32(maxX), float32(maxY), strokeWidth, color, true)
}
