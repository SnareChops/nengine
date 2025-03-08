package ui

import (
	"github.com/SnareChops/nengine/bit"
	"github.com/SnareChops/nengine/bounds"
	"github.com/SnareChops/nengine/types"
	"github.com/SnareChops/nengine/utils"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Button struct {
	*bounds.Raw
	state types.ButtonState
}

func NewButton(w, h int) types.Button {
	return &Button{Raw: new(bounds.Raw).Init(w, h)}
}

func (self *Button) Is(state types.ButtonState) bool {
	return bit.IsSet(self.state, state)
}

func (self *Button) ButtonState() types.ButtonState {
	return self.state
}

func (self *Button) Disable() {
	self.state |= types.ButtonStateDisabled
}

func (self *Button) Enable() {
	self.state &= ^types.ButtonStateDisabled
}

// Update updates the button. Requires the cursor x, y position
// in the same coordinate space as the button was positioned.
// Returns true if the state changed this frame
// Note: Use RelativePosition() if needed to convert global coordinates
// to relative coordinates.
func (self *Button) Update(x, y int) bool {
	prev := self.state
	self.state = 0
	if bit.IsSet(prev, types.ButtonStateDisabled) {
		self.state |= types.ButtonStateDisabled
	}
	if utils.IsWithin(self, x, y) {
		self.state |= types.ButtonStateHovered
		if !bit.IsSet(prev, types.ButtonStateHovered) {
			self.state |= types.ButtonStateJustHovered
		}
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			self.state |= types.ButtonStateClicked
			if !bit.IsSet(prev, types.ButtonStateClicked) && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
				self.state |= types.ButtonStateJustClicked
			}
		}
	}
	return prev != self.state
}
