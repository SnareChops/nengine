package ui

import (
	"strconv"

	"github.com/SnareChops/nengine/bounds"
	"github.com/SnareChops/nengine/input"
	"github.com/SnareChops/nengine/types"
	"github.com/SnareChops/nengine/utils"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type IntBox struct {
	*bounds.Raw
	keys      []ebiten.Key
	content   int
	cooldown  int
	repeating bool
	focused   bool
}

func (self *IntBox) Init(w, h int, input types.Input) *IntBox {
	self.Raw = new(bounds.Raw).Init(w, h)
	return self
}

func (self *IntBox) Content() int {
	return self.content
}

func (self *IntBox) SetContent(content int) {
	self.content = content
}

func (self *IntBox) Update(x, y, delta int) {
	if input.IsInputCaptured() {
		self.focused = false
		return
	}
	if self.focused {
		input.InputCapture()
		// Detect click outside of textbox to lose focus
		if !utils.IsWithin(self, x, y) && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			self.focused = false
			return
		}
		// Handle backspace
		self.cooldown -= delta
		if ebiten.IsKeyPressed(ebiten.KeyBackspace) {
			if self.cooldown <= 0 {
				self.content /= 10
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
			letter := ebiten.KeyName(key)
			value, err := strconv.Atoi(letter)
			if err != nil {
				continue
			}
			self.content = self.content*10 + value
		}
	} else {
		// Detect click on textbox to set focus
		if utils.IsWithin(self, x, y) && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			input.InputCapture()
			self.focused = true
		}
	}

}
