package emitters

import (
	"github.com/SnareChops/nengine/bounds"
	"github.com/SnareChops/nengine/types"
)

type Projectile struct {
	types.Bounds
	particles []Particle
}

func (self *Projectile) Init() *Projectile {
	self.Bounds = new(bounds.Raw).Init(0, 0)
	self.particles = []Particle{}
	return self
}

func (self *Projectile) Particles() []Particle {
	return self.particles
}

func (self *Projectile) Fire(particle Particle, duration int, angle, velocity float64) {
	particle.SetPos2(self.Pos2())
	particle.SetDuration(duration)
	particle.SetVelocity(angle, velocity)
	particle.Spawn()
	self.particles = append(self.particles, particle)
}

func (self *Projectile) Update(delta int) {
	i := 0
	for i < len(self.particles) {
		self.particles[i].Update(delta)
		if self.particles[i].Duration() <= 0 {
			self.particles[i].Despawn()
			self.particles = append(self.particles[:i], self.particles[i+1:]...)
		} else {
			i += 1
		}
	}
}
