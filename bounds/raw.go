package bounds

import (
	"github.com/SnareChops/nengine/types"
	"github.com/hajimehoshi/ebiten/v2"
)

// Raw A Bounds instance that represents an absolutely
// positioned bounds
type Raw struct {
	Position
	width    int
	height   int
	offsetX  float64
	offsetY  float64
	anchorX  int
	anchorY  int
	rotation float64
	scaleX   float64
	scaleY   float64
	options  *ebiten.DrawImageOptions
}

// Init sets the initial state of the RawBounds
func (self *Raw) Init(width, height int) *Raw {
	self.width = width
	self.height = height
	self.anchorX = LEFT
	self.anchorY = TOP
	self.scaleX = 1
	self.scaleY = 1
	self.options = &ebiten.DrawImageOptions{}
	return self
}

func (self *Raw) InitFromPoints(a types.Position, b types.Position) *Raw {
	x1 := min(a.X(), b.X())
	y1 := min(a.Y(), b.Y())
	x2 := max(a.X(), b.X())
	y2 := max(a.Y(), b.Y())
	self.Init(int(x2-x1), int(y2-y1))
	self.SetPos2(x1, y1)
	return self
}

// RawPos returns the raw position of the top left corner of the bounds as (x, y float64)
func (self *Raw) RawPos() (float64, float64) {
	ox, oy := self.Offset()
	return self.x - ox, self.y - oy
}

// SetAnchor sets the anchor point of the bounds to base it's position off
// Valid options for x: LEFT CENTER RIGHT
// Valid options for y: TOP CENTER BOTTOM
func (self *Raw) SetAnchor(x, y int) {
	self.anchorX, self.anchorY = x, y
	switch x {
	case LEFT:
		self.offsetX = 0
	case CENTER:
		self.offsetX = float64(self.width) / 2
	case RIGHT:
		self.offsetX = float64(self.width)
	}
	switch y {
	case TOP:
		self.offsetY = 0
	case CENTER:
		self.offsetY = float64(self.height) / 2
	case BOTTOM:
		self.offsetY = float64(self.height)
	}
}

// PosOf returns the pixel position of the specified anchor point
// Valid options for h: LEFT CENTER RIGHT
// Valid options for v: TOP CENTER BOTTOM
func (self *Raw) PosOf(h, v int) (x, y float64) {
	switch h {
	case LEFT:
		x = self.x - self.offsetX
	case CENTER:
		x = (self.x - self.offsetX) + (float64(self.width) / 2)
	case RIGHT:
		x = (self.x - self.offsetX) + float64(self.width)
	}

	switch v {
	case TOP:
		y = self.y - self.offsetY
	case CENTER:
		y = (self.y - self.offsetY) + (float64(self.height) / 2)
	case BOTTOM:
		y = (self.y - self.offsetY) + float64(self.height)
	}
	return
}

// Anchor return the current anchor position of the bounds (x, y int)
func (self *Raw) Anchor() (int, int) {
	return self.anchorX, self.anchorY
}

// Offset return the offset of the bounds based on it's anchor point
// This is the diff between the top left corner of the bounds and the anchor point
// (x, y float64)
func (self *Raw) Offset() (float64, float64) {
	return self.offsetX * self.scaleX, self.offsetY * self.scaleX
}

func (self *Raw) SetOffset(x, y float64) {
	self.offsetX = x
	self.offsetY = y
}

// Size returns the width and height of the bounds (width, height int)
func (self *Raw) Size() (int, int) {
	return int(float64(self.width) * self.scaleX), int(float64(self.height) * self.scaleY)
}

func (self *Raw) SetSize(w, h int) {
	self.width = w
	self.height = h
}

func (self *Raw) Resize(w, h int) {
	ow, oh := self.Size()
	self.width = w
	self.height = h
	self.offsetX = self.offsetX * (float64(w) / float64(ow))
	self.offsetY = self.offsetY * (float64(h) / float64(oh))
}

// Dx returns the width of the bounds
func (self *Raw) Dx() int {
	return int(float64(self.width) * self.scaleX)
}

// Dy returns the height of the bounds
func (self *Raw) Dy() int {
	return int(float64(self.height) * self.scaleY)
}

// SetRotation sets the rotation of the bounds around the center
func (self *Raw) SetRotation(theta float64) {
	self.rotation = theta
}

// Rotation return the rotation of the bounds
func (self *Raw) Rotation() float64 {
	return self.rotation
}

// Scale returns the scale of the bounds as (x, y float64)
func (self *Raw) Scale() (float64, float64) {
	return self.scaleX, self.scaleY
}

// SetScale sets the scale of the bounds
func (self *Raw) SetScale(x, y float64) {
	self.scaleX, self.scaleY = x, y
}

func (self *Raw) ScaleTo(w, h int) {
	wf := float64(w) / float64(self.width)
	hf := float64(h) / float64(self.height)
	self.scaleX = min(wf, hf)
	self.scaleY = min(wf, hf)
}

// Min returns the minimum coordinates of the bounds (x, y float64)
// This is the same as RawBounds, or the top left corner of the bounds
func (self *Raw) Min() (x, y float64) {
	return self.RawPos()
}

func (self *Raw) Mid() (x, y float64) {
	x, y = self.RawPos()
	w, h := self.Size()
	return x + float64(w)/2, y + float64(h)/2
}

// Max returns the maximum coordinates of the bounds (x, y float64)
// This is the same as the bottom right corner of the bounds
func (self *Raw) Max() (x, y float64) {
	x, y = self.RawPos()
	w, h := self.Size()
	return x + float64(w), y + float64(h)
}

func (self *Raw) MaxX() float64 {
	return self.x - self.offsetX + float64(self.width)*self.scaleX
}

func (self *Raw) MaxY() float64 {
	return self.y - self.offsetY + float64(self.height)*self.scaleY
}

// IsWithin checks if the provided x, y coordinate is within the bounds
func (self *Raw) IsWithin(x, y float64) bool {
	x1, y1 := self.RawPos()
	if self.width == 1 && self.height == 1 {
		return x == x1 && y == y1
	}
	w, h := self.Size()
	x2, y2 := x1+float64(w), y1+float64(h)
	return x > x1 && x < x2 && y > y1 && y < y2
}

// DoesCollide checks if the two bounds collide with each other
func (self *Raw) DoesCollide(other types.Bounds) bool {
	x1m, y1m := self.Min()
	x1M, y1M := self.Max()

	x2m, y2m := other.Min()
	x2M, y2M := other.Max()
	return !(x2M < x1m || x2m > x1M || y2M < y1m || y2m > y1M)
}

func (self *Raw) DrawOptions(camera types.Camera) *ebiten.DrawImageOptions {
	self.options.GeoM.Reset()
	rotation := self.Rotation()
	offx, offy := self.Offset()
	scalx, scaly := self.Scale()

	// Rotate around anchor
	if rotation != 0 {
		self.options.GeoM.Translate(-offx, -offy)
		self.options.GeoM.Rotate(rotation)
		self.options.GeoM.Translate(offx, offy)
	}
	// Scale
	if scalx != 1 && scaly != 1 {
		self.options.GeoM.Scale(scalx, scaly)
	}
	// Translate
	if camera == nil {
		self.options.GeoM.Translate(self.RawPos())
		return self.options
	}
	x, y := camera.WorldToScreenPos(self.RawPos())
	self.options.GeoM.Translate(float64(x), float64(y))
	return self.options
}
