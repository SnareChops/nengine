package nengine

import (
	"strings"

	"github.com/SnareChops/nengine/types"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type TextBox struct {
	*RawBounds
	input     *Input
	keys      []ebiten.Key
	content   string
	cooldown  int
	repeating bool
	focused   bool
}

func (self *TextBox) Init(w, h int, input types.Input) *TextBox {
	self.input = input.(*Input)
	self.RawBounds = new(RawBounds).Init(w, h)
	return self
}

func (self *TextBox) SetContent(content string) {
	self.content = content
}

func (self *TextBox) Content() string {
	return self.content
}

func (self *TextBox) Update(x, y, delta int) {
	if self.input.IsInputCaptured() {
		self.focused = false
		return
	}
	if self.focused {
		self.input.InputCapture()
		// Detect click outside of textbox to lose focus
		if !self.IsWithin(Floats(x, y)) && IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			self.focused = false
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
				letter = strings.ToUpper(letter)
			}
			self.content += letter
		}
	} else {
		// Detect click on textbox to set focus
		if self.IsWithin(Floats(x, y)) && IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			self.input.InputCapture()
			self.focused = true
		}
	}

}
