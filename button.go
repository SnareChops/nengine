package nengine

type ButtonState int

const (
	ButtonStateHovered ButtonState = 1 << iota
	ButtonStateClicked
	ButtonStateJustClicked
	ButtonStateJustHovered
)

type Button struct {
	*RawBounds
	state ButtonState
}

func (self *Button) Init(w, h int) *Button {
	self.RawBounds = new(RawBounds).Init(w, h)
	return self
}

func (self *Button) Is(state ButtonState) bool {
	return IsSet(self.state, state)
}

func (self *Button) ButtonState() ButtonState {
	return self.state
}

func (self *Button) Update(x, y int) {
	prev := self.state
	self.state = 0
	if self.IsWithin(Floats(x, y)) {
		self.state |= ButtonStateHovered
		if !IsSet(prev, ButtonStateHovered) {
			self.state |= ButtonStateJustHovered
		}
		if IsMouseButtonPressed(MouseButtonLeft) {
			self.state |= ButtonStateClicked
			if !IsSet(prev, ButtonStateClicked) {
				self.state |= ButtonStateJustClicked
			}
		}
	}
}
