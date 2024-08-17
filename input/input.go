package input

import (
	"github.com/SnareChops/nengine/types"
	"github.com/hajimehoshi/ebiten/v2"
)

type Input struct {
	captured bool

	cursorDeltaX  int
	cursorDeltaY  int
	cursorPrevX   int
	cursorPrevY   int
	cursorContent types.Sprite

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

func (self *Input) CursorContent() types.Sprite {
	return self.cursorContent
}

func (self *Input) SetCursorContent(content types.Sprite) {
	self.cursorContent = content
}

// CursorDelta returns the difference in screen
// cursor position from the previous frame
// Note: Update() must be called for this to function
// correctly
func (self *Input) CursorDelta() (int, int) {
	return self.cursorDeltaX, self.cursorDeltaY
}

func (self *Input) IsAnyMouseButtonPressed() bool {
	return ebiten.IsMouseButtonPressed(ebiten.MouseButton0) || ebiten.IsMouseButtonPressed(ebiten.MouseButton1) || ebiten.IsMouseButtonPressed(ebiten.MouseButton2)
}

func (self *Input) Update() {
	self.captured = false
	x, y := ebiten.CursorPosition()
	self.cursorDeltaX = x - self.cursorPrevX
	self.cursorDeltaY = y - self.cursorPrevY
	self.cursorPrevX = x
	self.cursorPrevY = y
	if self.cursorContent != nil {
		self.cursorContent.SetPos2(float64(x), float64(y))
	}
}

func (self *Input) Draw(screen *ebiten.Image) {
	if self.cursorContent != nil {
		screen.DrawImage(self.cursorContent.Image(), self.cursorContent.DrawOptions(nil))
	}
}
