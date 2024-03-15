package rendering

import (
	"image/color"

	"github.com/SnareChops/nengine/types"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func StrokeRect(dest *ebiten.Image, corner1, corner2 types.Position, width float32, color color.Color) {
	minX := min(corner1.X(), corner2.X())
	minY := min(corner1.Y(), corner2.Y())
	maxX := max(corner1.X(), corner2.X())
	maxY := max(corner1.Y(), corner2.Y())
	vector.StrokeRect(dest, float32(minX), float32(minY), float32(maxX-minX), float32(maxY-minY), width, color, false)
}
