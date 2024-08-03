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
	*Box
	velocity types.Vector
	options  *ebiten.DrawImageOptions
}

func (self *Projectile) Init(width, height int) *Projectile {
	self.Box = NewBox(width, height)
	return self
}

func (self *Projectile) Velocity() types.Vector {
	return self.velocity
}

func (self *Projectile) SetVelocity(velocity types.Vector) {
	self.velocity = velocity
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
