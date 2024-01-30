package bounds

import "github.com/SnareChops/nengine/types"

type VelocityBounds struct {
	*Raw
	velocity types.Vector
}

func (self *VelocityBounds) Init(width, height int) *VelocityBounds {
	self.Raw = new(Raw).Init(width, height)
	return self
}

func (self *VelocityBounds) Velocity() types.Vector {
	return self.velocity
}

func (self *VelocityBounds) SetVelocity(velocity types.Vector) {
	self.velocity = velocity
}

func (self *VelocityBounds) Update(delta int) {
	x, y := self.Pos2()
	self.SetPos2(x+self.velocity.X*float64(delta), y+self.velocity.Y*float64(delta))
}
