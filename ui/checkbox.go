package ui

import (
	"github.com/SnareChops/nengine/bit"
	"github.com/SnareChops/nengine/bounds"
	"github.com/SnareChops/nengine/utils"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type CheckboxState int

const (
	CheckboxStateChecked CheckboxState = 1 << iota
	CheckboxStateJustChanged
	CheckboxStateHovered
)

type Checkbox struct {
	*bounds.Raw
	state CheckboxState
}

func (self *Checkbox) Init(w, h int) *Checkbox {
	self.Raw = new(bounds.Raw).Init(w, h)
	return self
}

func (self *Checkbox) Is(state CheckboxState) bool {
	return bit.IsSet(self.state, state)
}

func (self *Checkbox) SetChecked(checked bool) {
	if checked {
		self.state = bit.BitmaskAdd(self.state, CheckboxStateChecked|CheckboxStateJustChanged)
	} else {
		self.state = bit.BitmaskRemove(self.state, CheckboxStateChecked|CheckboxStateJustChanged)
	}
}

func (self *Checkbox) Update(x, y int) bool {
	prev := self.state
	self.state = bit.BitmaskRemove(self.state, CheckboxStateHovered|CheckboxStateJustChanged)
	if utils.IsWithin(self, x, y) {
		self.state = bit.BitmaskAdd(self.state, CheckboxStateHovered)
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			self.state = bit.BitmaskAdd(self.state, CheckboxStateJustChanged)
			if bit.IsSet(self.state, CheckboxStateChecked) {
				self.state = bit.BitmaskRemove(self.state, CheckboxStateChecked)
			} else {
				self.state = bit.BitmaskAdd(self.state, CheckboxStateChecked)
			}
		}
	}
	return prev != self.state
}
