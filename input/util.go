package input

import (
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
)

func KeyToUpper(key ebiten.Key, letter string) string {
	switch key {
	case ebiten.KeyBackquote:
		return "~"
	case ebiten.KeyMinus:
		return "_"
	case ebiten.KeyEqual:
		return "+"
	case ebiten.KeyLeftBracket:
		return "{"
	case ebiten.KeyRightBracket:
		return "}"
	case ebiten.KeyBackslash:
		return "|"
	case ebiten.KeySemicolon:
		return ":"
	case ebiten.KeyApostrophe:
		return "\""
	case ebiten.KeyComma:
		return "<"
	case ebiten.KeyPeriod:
		return ">"
	case ebiten.KeySlash:
		return "?"
	case ebiten.Key1:
		return "!"
	case ebiten.Key2:
		return "@"
	case ebiten.Key3:
		return "#"
	case ebiten.Key4:
		return "$"
	case ebiten.Key5:
		return "%"
	case ebiten.Key6:
		return "^"
	case ebiten.Key7:
		return "&"
	case ebiten.Key8:
		return "*"
	case ebiten.Key9:
		return "("
	case ebiten.Key0:
		return ")"
	}
	return strings.ToUpper(letter)
}
