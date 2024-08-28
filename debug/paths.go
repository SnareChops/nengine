package debug

import (
	"image/color"
	"unsafe"

	"github.com/SnareChops/nengine/types"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type path struct {
	points []types.Position
	color  color.Color
}

var paths map[unsafe.Pointer]path

func AddPath(ptr unsafe.Pointer, points []types.Position, color color.Color) {
	if _, ok := paths[ptr]; ok {
		return
	}
	paths[ptr] = path{points, color}
}

func RemovePath(ptr unsafe.Pointer) {
	delete(paths, ptr)
}

func DrawPaths(screen *ebiten.Image, camera types.Camera) {
	for _, path := range paths {
		for i, j := 0, 1; j < len(path.points); i, j = i+1, j+1 {
			x1, y1 := camera.WorldToScreenPos(path.points[i].Pos2())
			x2, y2 := camera.WorldToScreenPos(path.points[j].Pos2())
			vector.StrokeLine(screen, float32(x1), float32(y1), float32(x2), float32(y2), 5, path.color, true)
		}
	}
}
