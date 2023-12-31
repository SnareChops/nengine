package bounds

import "math"

type Physics struct {
	Raw
	vx float64
	vy float64
	ax float64
	ay float64
}

func (self *Physics) Init(width, height int) *Physics {
	self.Raw.Init(width, height)
	return self
}

// Velocity Gets the horizontal and vertical velocity of the bounds
// returns (h, v float64)
func (self *Physics) Velocity() (float64, float64) {
	return self.vx, self.vy
}

// SetRawVelocity Sets the horizontal and vertical velocity of the bounds
func (self *Physics) SetRawVelocity(x, y float64) {
	self.vx = x
	self.vy = y
}

// SetVelocity Sets the angle and magnitude of the bounds velocity
func (self *Physics) SetVelocity(angle, magnitude float64) {
	self.vx = math.Cos(angle) * magnitude
	self.vy = math.Sin(angle) * magnitude
}

// Acceleration Gets the acceration of the bounds
// returns (x, y float64)
func (self *Physics) Acceleration() (float64, float64) {
	return self.ax, self.ay
}

func (self *Physics) SetAcceleration(angle, magnitude float64) {
	self.ax = math.Cos(angle) * magnitude
	self.ay = math.Sin(angle) * magnitude
}

func (self *Physics) Update(delta int) {
	self.vx += self.ax * float64(delta)
	self.vy += self.ay * float64(delta)
	x, y := self.Pos2()
	self.SetPos2(x+self.vx*float64(delta), y+self.vy*float64(delta))
}
