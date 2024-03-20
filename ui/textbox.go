package ui

import (
	"strings"

	"github.com/SnareChops/nengine/bounds"
	_input "github.com/SnareChops/nengine/input"
	"github.com/SnareChops/nengine/types"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type TextBox struct {
	*bounds.Raw
	input     *_input.Input
	keys      []ebiten.Key
	content   string
	cooldown  int
	repeating bool
	focused   bool
}

func (self *TextBox) Init(w, h int, input types.Input) *TextBox {
	self.input = input.(*_input.Input)
	self.Raw = new(bounds.Raw).Init(w, h)
	return self
}

func (self *TextBox) SetContent(content string) {
	self.content = content
}

func (self *TextBox) Content() string {
	return self.content
}

func (self *TextBox) Focus() {
	self.focused = true
}

func (self *TextBox) IsFocused() bool {
	return self.focused
}

func (self *TextBox) Update(x, y, delta int) {
	if self.input.IsInputCaptured() {
		self.focused = false
		return
	}
	if self.focused {
		self.input.InputCapture()
		// Detect click outside of textbox to lose focus
		if !self.IsWithin(float64(x), float64(y)) && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			self.focused = false
			self.input.InputUncapture()
			return
		}
		// Handle backspace
		self.cooldown -= delta
		if ebiten.IsKeyPressed(ebiten.KeyBackspace) {
			if self.cooldown <= 0 {
				self.content = self.content[:len(self.content)-1]
				self.repeating = true
				if self.repeating {
					self.cooldown = 100
				} else {
					self.cooldown = 500
				}
			}
			return
		} else {
			self.cooldown = 0
			self.repeating = false
		}
		// Handle keypresses
		self.keys = inpututil.AppendJustPressedKeys(self.keys[:0])
		for _, key := range self.keys {
			if key == ebiten.KeySpace {
				self.content += " "
				continue
			}
			letter := ebiten.KeyName(key)
			if ebiten.IsKeyPressed(ebiten.KeyShift) {
				letter = keyToUpper(key, letter)
				letter = strings.ToUpper(letter)
			}
			self.content += letter
		}
	} else {
		// Detect click on textbox to set focus
		if self.IsWithin(float64(x), float64(y)) && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			self.input.InputCapture()
			self.focused = true
		}
	}
}

func keyToUpper(key ebiten.Key, letter string) string {
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
