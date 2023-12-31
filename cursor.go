package nengine

// type ICursor interface {
// 	CursorUpdate()
// 	CursorContent() Bounds
// 	SetCursorContent(Bounds)
// 	CursorDelta() (int, int)
// 	CursorCapture()
// 	CursorCaptured() bool
// }

// type Cursor struct {
// 	captured     bool
// 	prevScreenX  int
// 	prevScreenY  int
// 	deltaScreenX int
// 	deltaScreenY int
// 	content      Bounds
// }

// func (self *Cursor) CursorUpdate() {
// 	self.captured = false
// 	x, y := CursorPosition()
// 	self.deltaScreenX = x - self.prevScreenX
// 	self.deltaScreenY = y - self.prevScreenY
// 	self.prevScreenX = x
// 	self.prevScreenY = y
// 	if self.content != nil {
// 		self.content.SetPos2(Floats(x, y))
// 	}
// }

// func (self *Cursor) CursorContent() Bounds {
// 	return self.content
// }

// func (self *Cursor) SetCursorContent(content Bounds) {
// 	self.content = content
// }

// // CursorDelta returns the difference in screen
// // cursor position from the previous frame
// // Note: Update() must be called for this to function
// // correctly
// func (self *Cursor) CursorDelta() (int, int) {
// 	return self.deltaScreenX, self.deltaScreenY
// }

// func (self *Cursor) CursorCapture() {
// 	self.captured = true
// }

// func (self *Cursor) CursorCaptured() bool {
// 	return self.captured
// }
