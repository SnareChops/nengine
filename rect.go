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
	for a := 0; a < len(rects); a++ {
		b := a + 1
		for b < len(rects) {
			// If a is the same height as b
			// and a is at the same y value as b
			// and b is touching the left side of a
			if rects[a].MinY == rects[b].MinY && rects[a].MaxY == rects[b].MaxY && rects[a].MaxX == rects[b].MinX {
				// Combine b into a and delete b
				rects[a].MaxX = rects[b].MaxX
				rects = slices.Delete(rects, b, b+1)
			} else {
				b++
			}
		}
	}

	// Combine vertically
	for a := 0; a < len(rects); a++ {
		b := a + 1
		for b < len(rects) {
			// If a is same width as b
			// and a is at same x value as b
			// and b is touching the bottom of a
			if rects[a].MinX == rects[b].MinX && rects[a].MaxX == rects[b].MaxX && rects[a].MaxY == rects[b].MinY {
				// Combine b into a and delete b
				rects[a].MaxY = rects[b].MaxY
				rects = slices.Delete(rects, b, b+1)
			} else {
				b++
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
