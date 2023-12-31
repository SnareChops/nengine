package nengine

import (
	"github.com/SnareChops/nengine/bounds"
	"github.com/SnareChops/nengine/types"
	"golang.org/x/exp/slices"
)

type Rect struct {
	MinX int
	MinY int
	MaxX int
	MaxY int
}

func CombineRects(rects []Rect) []Rect {
	// Combine Horizontally
	for _, a := range rects {
		for i, b := range rects {
			// Skip a == b
			if a == b {
				continue
			}

			// If a is the same height as b
			// and a is at the same y value as b
			// and b is touching the left side of a
			if a.MinY == b.MinY && a.MaxY == b.MaxY && a.MaxX == b.MinX {
				// Combine b into a and delete b
				a.MaxX = b.MaxX
				rects = slices.Delete(rects, i, i+1)
			}
		}
	}

	// Combine vertically
	for _, a := range rects {
		for i, b := range rects {
			// Skip a == b
			if a == b {
				continue
			}
			// If a is same width as b
			// and a is at same x value as b
			// and b is touching the bottom of a
			if a.MinX == b.MinX && a.MaxX == b.MaxX && a.MaxY == b.MinY {
				// Combine b into a and delete b
				a.MaxY = b.MaxY
				rects = slices.Delete(rects, i, i+1)
			}
		}
	}
	return rects
}

func RectsToBounds(rects []Rect) []types.Bounds {
	out := []types.Bounds{}
	for _, rect := range rects {
		bounds := new(bounds.Raw).Init(rect.MaxX-rect.MinX, rect.MaxY-rect.MinY)
		x, y := Floats(rect.MinX, rect.MinY)
		bounds.SetPos2(x, y)
		out = append(out, bounds)
	}
	return out
}
