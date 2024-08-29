package input

import (
	"github.com/SnareChops/nengine/types"
	"github.com/hajimehoshi/ebiten/v2"
)

type input struct {
	captured      bool
	cursorDeltaX  int
	cursorDeltaY  int
	cursorPrevX   int
	cursorPrevY   int
	cursorContent types.Sprite
}

var state = &input{}

func InputCapture() {
	state.captured = true
}

func InputUncapture() {
	state.captured = false
}

func IsInputCaptured() bool {
	return state.captured
}

func CursorContent() types.Sprite {
	return state.cursorContent
}

func SetCursorContent(content types.Sprite) {
	state.cursorContent = content
}

func HideCursor() {
	ebiten.SetCursorMode(ebiten.CursorModeHidden)
}

func ShowCursor() {
	ebiten.SetCursorMode(ebiten.CursorModeVisible)
}

// CursorDelta returns the difference in screen
// cursor position from the previous frame
func CursorDelta() (int, int) {
	return state.cursorDeltaX, state.cursorDeltaY
}

func IsAnyMouseButtonPressed() bool {
	return ebiten.IsMouseButtonPressed(ebiten.MouseButton0) || ebiten.IsMouseButtonPressed(ebiten.MouseButton1) || ebiten.IsMouseButtonPressed(ebiten.MouseButton2)
}

func Reset() {
	state = &input{}
}

func Update() {
	state.captured = false
	x, y := ebiten.CursorPosition()
	state.cursorDeltaX = x - state.cursorPrevX
	state.cursorDeltaY = y - state.cursorPrevY
	state.cursorPrevX = x
	state.cursorPrevY = y
	if state.cursorContent != nil {
		state.cursorContent.SetPos2(float64(x), float64(y))
	}
}

func Draw(screen *ebiten.Image) {
	if state.cursorContent != nil {
		screen.DrawImage(state.cursorContent.Image(), state.cursorContent.DrawOptions(nil))
	}
}
