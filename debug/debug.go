package debug

import (
	"image/color"
	"unsafe"

	"github.com/SnareChops/nengine/types"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

var (
	debugEnabled bool
	debugFont    font.Face
	debugStats   []stat = []stat{}
	debugPaths   map[unsafe.Pointer]path
)

type stat struct {
	label string
	value func() string
}

func EnableDebug(face font.Face) {
	debugEnabled = true
	debugFont = face
}

func DebugEnabled() bool {
	return debugEnabled
}

func DebugStat(label string, value func() string) {
	for _, stat := range debugStats {
		if stat.label == label {
			stat.value = value
			return
		}
	}
	debugStats = append(debugStats, stat{label, value})
}

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

// Draw the debug information to the provided image (usually the screen)
func DebugDraw(screen *ebiten.Image) {
	if !debugEnabled {
		return
	}
	s := ""
	for _, stat := range debugStats {
		if stat.value != nil {
			s += stat.label + ": " + stat.value() + "\n"
		}
	}
	text.Draw(screen, s, debugFont, 10, 70, color.White)
}
