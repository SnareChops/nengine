package emitters

import (
	"github.com/SnareChops/nengine/bounds"
	"github.com/SnareChops/nengine/types"
	"github.com/SnareChops/nengine/utils"
)

type Uniform struct {
	types.Bounds
	particles []Particle
	active    []Particle
	velocity  float64
	minAngle  float64
	maxAngle  float64
	life      int
	density   int
}

func (self *Uniform) Init(particles []Particle) *Uniform {
	self.Bounds = new(bounds.Raw).Init(0, 0)
	self.particles = particles
	self.active = make([]Particle, len(particles))
	return self
}

func (self *Uniform) SetVelocity(velocity float64) {
	self.velocity = velocity
}

// SetAngle sets the minimum and maximum emittion angles in radians
func (self *Uniform) SetAngle(min, max float64) {
	self.minAngle = min
	self.maxAngle = max
}

func (self *Uniform) SetLife(life int) {
	self.life = life
}

func (self *Uniform) SetDensity(density int) {
	self.density = density
}

func (self *Uniform) Particles() []Particle {
	return self.active
}

func (self *Uniform) Start() {
	for i := 0; i < self.density; i++ {
		for _, particle := range self.particles {
			if particle.Duration() > 0 {
				continue
			}
			percent := float64(i) / float64(self.density)
			angle := utils.LinearInterpolate(self.minAngle, self.maxAngle, percent)
			particle.SetPos2(self.Pos2())
			particle.SetDuration(self.life)
			particle.SetVelocity(angle, self.velocity)
			particle.Spawn()
			for i, active := range self.active {
				if active == nil {
					self.active[i] = particle
					break
				}
			}
			break
		}
	}
}

func (self *Uniform) Update(delta int) {
	if len(self.active) == 0 {
		return
	}
	i := 0
	for i < len(self.active) {
		self.active[i].Update(delta)
		if self.active[i].Duration() <= 0 {
			self.active[i].Despawn()
			self.active = append(self.active[:i], self.active[i+1:]...)
		} else {
			i += 1
		}
	}
}
