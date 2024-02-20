package bounds

import (
	"github.com/SnareChops/nengine/types"
	"github.com/hajimehoshi/ebiten/v2"
)

func Point(x, y, z float64) *Position {
	return &Position{x, y, z, &ebiten.DrawImageOptions{}}
}

type Position struct {
	x       float64
	y       float64
	z       float64
	options *ebiten.DrawImageOptions
}

func (self *Position) Pos() types.Vector {
	return types.Vector{self.x, self.y}
}

func (self *Position) SetPos(pos types.Vector) {
	self.x = pos.X
	self.y = pos.Y
}

// Pos2 returns the x and y components of the Vector (x, y float64)
func (self *Position) Pos2() (float64, float64) {
	return self.x, self.y
}

// Pos3 returns all components of the Vector (x, y, z float64)
func (self *Position) Pos3() (float64, float64, float64) {
	return self.x, self.y, self.z
}

// SetPos2 sets the x and y components of the Vector
func (self *Position) SetPos2(x, y float64) {
	self.x = x
	self.y = y
}

// SetPos3 sets the x, y, and z components of the Vector
func (self *Position) SetPos3(x, y, z float64) {
	self.x = x
	self.y = y
	self.z = z
}

func (self *Position) X() float64 {
	return self.x
}

func (self *Position) Y() float64 {
	return self.y
}

func (self *Position) Z() float64 {
	return self.z
}

func (self *Position) GridAlign(h, v int) {
	self.x = float64(int(self.x) / h * h)
	self.y = float64(int(self.y) / v * v)
}

func (self *Position) DrawOptions(camera types.Camera) *ebiten.DrawImageOptions {
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
