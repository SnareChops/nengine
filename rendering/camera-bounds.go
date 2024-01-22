package rendering

type CameraBounds struct {
	x, y   float64
	w, h   int
	ox, oy float64
}

func (self *CameraBounds) Init(w, h int) *CameraBounds {
	self.w, self.h = w, h
	self.ox, self.oy = float64(w)/2, float64(h)/2
	return self
}

func (self *CameraBounds) Pos() (float64, float64) {
	return self.x, self.y
}

func (self *CameraBounds) SetPos(x, y float64) {
	self.x, self.y = x, y
}

func (self *CameraBounds) Size() (int, int) {
	return self.w, self.h
}

func (self *CameraBounds) Resize(w, h int) {
	ow, oh := self.Size()
	self.w, self.h = w, h
	self.ox = self.ox * (float64(w) / float64(ow))
	self.oy = self.oy * (float64(h) / float64(oh))
}

func (self *CameraBounds) Dx() int {
	return self.w
}

func (self *CameraBounds) Dy() int {
	return self.h
}

func (self *CameraBounds) Min() (float64, float64) {
	return self.x - self.ox, self.y - self.oy
}

func (self *CameraBounds) Max() (float64, float64) {
	x, y := self.Min()
	w, h := self.Size()
	return x + float64(w), y + float64(h)
}

func (self *CameraBounds) X() float64 {
	return self.x
}

func (self *CameraBounds) Y() float64 {
	return self.y
}

func (self *CameraBounds) MinX() float64 {
	return self.x - self.ox
}

func (self *CameraBounds) MinY() float64 {
	return self.y - self.oy
}

func (self *CameraBounds) MaxX() float64 {
	return self.x - self.ox + float64(self.Dx())
}

func (self *CameraBounds) MaxY() float64 {
	return self.y - self.oy + float64(self.Dy())
}
