package nengine

import "github.com/hajimehoshi/ebiten/v2"

type Input struct {
	captured bool

	cursorDeltaX  int
	cursorDeltaY  int
	cursorPrevX   int
	cursorPrevY   int
	cursorContent Sprite

	order int
}

func (self *Input) SetOrder(order int) {
	self.order = order
}

func (self *Input) Order() int {
	return self.order
}

func (self *Input) InputCapture() {
	self.captured = true
}

func (self *Input) InputUncapture() {
	self.captured = false
}

func (self *Input) IsInputCaptured() bool {
	return self.captured
}

func (self *Input) CursorContent() Sprite {
	return self.cursorContent
}

func (self *Input) SetCursorContent(content Sprite) {
	self.cursorContent = content
}

// CursorDelta returns the difference in screen
// cursor position from the previous frame
// Note: Update() must be called for this to function
// correctly
func (self *Input) CursorDelta() (int, int) {
	return self.cursorDeltaX, self.cursorDeltaY
}

func (self *Input) Update() {
	self.captured = false
	x, y := CursorPosition()
	self.cursorDeltaX = x - self.cursorPrevX
	self.cursorDeltaY = y - self.cursorPrevY
	self.cursorPrevX = x
	self.cursorPrevY = y
	if self.cursorContent != nil {
		self.cursorContent.SetPos2(Floats(x, y))
	}
}

func (self *Input) Draw(screen *ebiten.Image) {
	if self.cursorContent != nil {
		screen.DrawImage(self.cursorContent.Image(), self.cursorContent.DrawOptions(nil))
	}
}
