package ui

import (
	"strings"

	"github.com/SnareChops/nengine/bounds"
	"github.com/SnareChops/nengine/input"
	"github.com/SnareChops/nengine/utils"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type TextBox struct {
	*bounds.Raw
	keys      []ebiten.Key
	content   string
	cooldown  int
	repeating bool
	focused   bool
}

func (self *TextBox) Init(w, h int) *TextBox {
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
	if input.IsInputCaptured() {
		self.focused = false
		return
	}
	if self.focused {
		input.InputCapture()
		// Detect click outside of textbox to lose focus
		if !utils.IsWithin(self, x, y) && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			self.focused = false
			input.InputUncapture()
			return
		}
		// Handle backspace
		self.cooldown -= delta
		if ebiten.IsKeyPressed(ebiten.KeyBackspace) {
			if self.cooldown <= 0 {
				self.content = self.content[:len(self.content)-1]
				if self.repeating {
					self.cooldown = 100
				} else {
					self.cooldown = 500
				}
				self.repeating = true
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
				letter = input.KeyToUpper(key, letter)
				letter = strings.ToUpper(letter)
			}
			self.content += letter
		}
	} else {
		// Detect click on textbox to set focus
		if utils.IsWithin(self, x, y) && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			input.InputCapture()
			self.focused = true
		}
	}
}
