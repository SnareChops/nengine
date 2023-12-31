package nengine

import (
	"io"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type ColorScale = ebiten.ColorScale

var (
	// OS
	ReadFile = os.ReadFile
	Open     = func(name string) (io.Reader, error) {
		return os.Open(name)
	}

	// Mouse
	CursorPosition = func() (int, int) {
		return ebiten.CursorPosition()
	}
	MouseButtonLeft           = ebiten.MouseButton0
	MouseButtonMiddle         = ebiten.MouseButton1
	MouseButtonRight          = ebiten.MouseButton2
	MouseButtonBack           = ebiten.MouseButton3
	MouseButtonForward        = ebiten.MouseButton4
	IsMouseButtonPressed      = ebiten.IsMouseButtonPressed
	IsMouseButtonJustPressed  = inpututil.IsMouseButtonJustPressed
	IsMouseButtonJustReleased = inpututil.IsMouseButtonJustReleased
	Wheel                     = ebiten.Wheel

	// Keyboard
	IsKeyPressed     = ebiten.IsKeyPressed
	IsKeyJustPressed = inpututil.IsKeyJustPressed
)
