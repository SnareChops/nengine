package debug

import (
	"image/color"
	"unsafe"

	"github.com/SnareChops/nengine/types"
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
func Draw(screen types.Image) {
	y := 70
	h := debugFont.Metrics().Height.Ceil()
	if enableStats {
		for _, stat := range stats {
			if stat.value != nil {
				text.Draw(screen.Raw(), stat.label+": "+stat.value(), debugFont, 10, y, color.White)
				y += h
			}
		}
	}
	if enableTimers {
		for _, timer := range timers {
			text.Draw(screen.Raw(), timer.name+": "+timer.Value(), debugFont, 10, y, color.White)
			y += h
		}
	}
	if enableTimers {
		for _, timer := range FrameTimers {
			text.Draw(screen.Raw(), timer.name+": "+timer.Value(), debugFont, 10, y, color.White)
			y += h
		}
	}
}
