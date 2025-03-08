package ui

import (
	"github.com/SnareChops/nengine/bit"
	"github.com/SnareChops/nengine/bounds"
	"github.com/SnareChops/nengine/types"
	"github.com/SnareChops/nengine/utils"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Checkbox struct {
	*bounds.Raw
	state types.CheckboxState
}

func NewCheckbox(w, h int) types.Checkbox {
	return &Checkbox{Raw: new(bounds.Raw).Init(w, h)}
}

func (self *Checkbox) Is(state types.CheckboxState) bool {
	return bit.IsSet(self.state, state)
}

func (self *Checkbox) SetChecked(checked bool) {
	if checked {
		self.state = bit.BitmaskAdd(self.state, types.CheckboxStateChecked|types.CheckboxStateJustChanged)
	} else {
		self.state = bit.BitmaskRemove(self.state, types.CheckboxStateChecked|types.CheckboxStateJustChanged)
	}
}

func (self *Checkbox) Update(x, y int) bool {
	prev := self.state
	self.state = bit.BitmaskRemove(self.state, types.CheckboxStateHovered|types.CheckboxStateJustChanged)
	if utils.IsWithin(self, x, y) {
		self.state = bit.BitmaskAdd(self.state, types.CheckboxStateHovered)
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			self.state = bit.BitmaskAdd(self.state, types.CheckboxStateJustChanged)
			if bit.IsSet(self.state, types.CheckboxStateChecked) {
				self.state = bit.BitmaskRemove(self.state, types.CheckboxStateChecked)
			} else {
				self.state = bit.BitmaskAdd(self.state, types.CheckboxStateChecked)
			}
		}
	}
	return prev != self.state
}
