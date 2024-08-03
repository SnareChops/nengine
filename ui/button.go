package ui

import (
	"github.com/SnareChops/nengine/bit"
	"github.com/SnareChops/nengine/bounds"
	"github.com/SnareChops/nengine/utils"
	"github.com/hajimehoshi/ebiten/v2"
)

type ButtonState int

const (
	ButtonStateHovered ButtonState = 1 << iota
	ButtonStateClicked
	ButtonStateJustClicked
	ButtonStateJustHovered
	ButtonStateDisabled
)

type Button struct {
	*bounds.Raw
	state ButtonState
}

func (self *Button) Init(w, h int) *Button {
	self.Raw = new(bounds.Raw).Init(w, h)
	return self
}

func (self *Button) Is(state ButtonState) bool {
	return bit.IsSet(self.state, state)
}

func (self *Button) ButtonState() ButtonState {
	return self.state
}

func (self *Button) Disable() {
	self.state |= ButtonStateDisabled
}

func (self *Button) Enable() {
	self.state &= ^ButtonStateDisabled
}

// Update updates the button. Requires the cursor x, y position
// in the same coordinate space as the button was positioned.
// Returns true if the state changed this frame
// Note: Use RelativePosition() if needed to convert global coordinates
// to relative coordinates.
func (self *Button) Update(x, y int) bool {
	prev := self.state
	self.state = 0
	if bit.IsSet(prev, ButtonStateDisabled) {
		self.state |= ButtonStateDisabled
	}
	if utils.IsWithin(self, x, y) {
		self.state |= ButtonStateHovered
		if !bit.IsSet(prev, ButtonStateHovered) {
			self.state |= ButtonStateJustHovered
		}
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			self.state |= ButtonStateClicked
			if !bit.IsSet(prev, ButtonStateClicked) {
				self.state |= ButtonStateJustClicked
			}
		}
	}
	return prev != self.state
}
