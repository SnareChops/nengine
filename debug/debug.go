package debug

import (
	"image/color"
	"unsafe"

	"github.com/SnareChops/nengine/types"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

var (
	debugPaths map[unsafe.Pointer]path
)

func DebugPath(ptr unsafe.Pointer, points []types.Position, color color.Color) {
	for key := range debugPaths {
		if key == ptr {
			debugPaths[ptr] = path{points, color}
			return
		}
	}
	debugPaths[ptr] = path{points, color}
}

func Paths() map[unsafe.Pointer]path {
	return debugPaths
}

func Update() {
	for _, timer := range FrameTimers {
		timer.EndFrame()
	}
}

// Draw the debug information to the provided image (usually the screen)
func Draw(screen *ebiten.Image) {
	s := ""
	for _, stat := range stats {
		if stat.value != nil {
			s += stat.label + ": " + stat.value() + "\n"
		}
	}
	text.Draw(screen, s, debugFont, 10, 70, color.White)
}
