package bounds

type Box struct {
	*Position
	w, h     int
	ox, oy   float64
	rotation float64
}

func NewBox(w, h int) *Box {
	return &Box{Position: new(Position), w: w, h: h}
}

func (self *Box) Size() (int, int) {
	return self.w, self.h
}

func (self *Box) SetSize(w, h int) {
	self.w, self.h = w, h
}

func (self *Box) Resize(w, h int) {
	ow, oh := self.Size()
	self.w, self.h = w, h
	self.ox = self.ox * (float64(w) / float64(ow))
	self.oy = self.oy * (float64(h) / float64(oh))
}

func (self *Box) Offset() (float64, float64) {
	return self.ox, self.oy
}

func (self *Box) SetOffset(x, y float64) {
	self.ox, self.oy = x, y
}

func (self *Box) Rotation() float64 {
	return self.rotation
}

func (self *Box) SetRotation(theta float64) {
	self.rotation = theta
}

// SetAnchor sets the anchor point of the bounds to base it's position off
// Valid options for x: LEFT CENTER RIGHT
// Valid options for y: TOP CENTER BOTTOM
func (self *Box) SetAnchor(x, y int) {
	switch x {
	case LEFT:
		self.ox = 0
	case CENTER:
		self.ox = float64(self.Dx()) / 2
	case RIGHT:
		self.ox = float64(self.Dx()) - 1
	}
	switch y {
	case TOP:
		self.oy = 0
	case CENTER:
		self.oy = float64(self.Dy()) / 2
	case BOTTOM:
		self.oy = float64(self.Dy()) - 1
	}
}

func (self *Box) Dx() int {
	return self.w
}

func (self *Box) Dy() int {
	return self.h
}

func (self *Box) Min() (x, y float64) {
	x = self.x - self.ox
	y = self.y - self.oy
	return
}

func (self *Box) Mid() (x, y float64) {
	x = (float64(self.w) / 2) + self.x - self.ox
	y = (float64(self.h) / 2) + self.y - self.oy
	return
}

func (self *Box) Max() (x, y float64) {
	x = self.x - self.ox + float64(self.w) - 1
	y = self.y - self.oy + float64(self.h) - 1
	return
}

func (self *Box) MidX() float64 {
	return (float64(self.w) / 2) + self.x - self.ox
}

func (self *Box) MidY() float64 {
	return (float64(self.h) / 2) + self.y - self.oy
}

func (self *Box) MaxX() float64 {
	return self.x - self.ox + float64(self.w) - 1
}

func (self *Box) MaxY() float64 {
	return self.y - self.oy + float64(self.h) - 1
}
