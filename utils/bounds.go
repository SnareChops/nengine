package utils

import (
	"github.com/SnareChops/nengine/bounds"
	"github.com/SnareChops/nengine/types"
)

// ChunkBounds returns a list of bounds split from the original bounds
// using the maximum provided size
// Note: This is usually used to facilitate splitting a large Bounds or image
// into smaller pieces so that it can be used by the Renderer
func ChunkBounds(original types.Bounds, size int) []types.Bounds {
	out := []types.Bounds{}
	x, y := original.Pos2()
	width, height := original.Size()
	for i := 0; i < width; i += size {
		for j := 0; j < height; j += size {
			var w int
			var h int
			if i+size > width {
				w = width - i
			} else {
				w = size
			}

			if j+size > height {
				h = height - j
			} else {
				h = size
			}

			bounds := new(bounds.Raw).Init(w, h)
			bounds.SetPos2(x+float64(i), y+float64(j))
			out = append(out, bounds)
		}
	}
	return out
}
