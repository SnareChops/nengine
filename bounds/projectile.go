package bounds

import (
	"github.com/SnareChops/nengine/types"
	"github.com/hajimehoshi/ebiten/v2"
)

// A slimmed down projectile bounds with the minimum amount
// of memory usage for supporting many simple projectiles
// Should work with most things that you need for projectiles
// but by design is a subset of Bounds functionality
type Projectile struct {
	x, y     float64
	w, h     int
	velocity types.Vector
	options  *ebiten.DrawImageOptions
}

func (self *Projectile) Init(width, height int) *Projectile {
	self.w = width
	self.h = height
	return self
}

func (self *Projectile) Velocity() types.Vector {
	return self.velocity
}

func (self *Projectile) SetVelocity(velocity types.Vector) {
	self.velocity = velocity
}

func (self *Projectile) Pos2() (float64, float64) {
	return self.x, self.y
}

func (self *Projectile) SetPos2(x, y float64) {
	self.x = x
	self.y = y
}

func (self *Projectile) RawPos() (float64, float64) {
	return self.x, self.y
}

func (self *Projectile) X() float64 {
	return self.x
}

func (self *Projectile) Y() float64 {
	return self.y
}

func (self *Projectile) Size() (int, int) {
	return self.w, self.h
}

func (self *Projectile) Dx() int {
	return self.w
}

func (self *Projectile) Dy() int {
	return self.h
}

func (self *Projectile) Min() (x, y float64) {
	return self.RawPos()
}

func (self *Projectile) Max() (x, y float64) {
	x, y = self.RawPos()
	w, h := self.Size()
	return x + float64(w-1), y + float64(h-1)
}

func (self *Projectile) IsWithin(x, y float64) bool {
	x1, y1 := self.RawPos()
	if self.w == 1 && self.h == 1 {
		return x == x1 && y == y1
	}
	x2 := x1 + float64(self.w-1)
	y2 := y1 + float64(self.h-1)
	return x >= x1 && x <= x2 && y >= y1 && y <= y2
}

func (self *Projectile) DoesCollide(other types.Collidable) bool {
	x1m, y1m := self.Min()
	x1M, y1M := self.Max()

	x2m, y2m := other.Min()
	x2M, y2M := other.Max()
	return !(x2M < x1m || x2m > x1M || y2M < y1m || y2m > y1M)
}

func (self *Projectile) Update(delta int) {
	x, y := self.Pos2()
	self.SetPos2(x+self.velocity.X*float64(delta), y+self.velocity.Y*float64(delta))
}

func (self *Projectile) DrawOptions(camera types.Camera) *ebiten.DrawImageOptions {
	if self.options == nil {
		self.options = &ebiten.DrawImageOptions{}
	}
	self.options.GeoM.Reset()
	if camera == nil {
		self.options.GeoM.Translate(self.x, self.y)
		return self.options
	}
	x, y := camera.WorldToScreenPos(self.x, self.y)
	self.options.GeoM.Translate(float64(x), float64(y))
	return self.options
}
