package emitters

import (
	"github.com/SnareChops/nengine/bounds"
)

type Particle interface {
	bounds.PhysicsBounds
	Update(delta int)
	Duration() int
	SetDuration(ms int)
	Spawn()
	Despawn()
}

type ParticleBase struct {
	bounds.PhysicsBounds
	duration int
}

func (self *ParticleBase) Init(width, height int) *ParticleBase {
	self.PhysicsBounds = new(bounds.Physics).Init(width, height)
	return self
}

func (self *ParticleBase) Duration() int {
	return self.duration
}

func (self *ParticleBase) SetDuration(ms int) {
	self.duration = ms
}

func (self *ParticleBase) Update(delta int) {
	self.PhysicsBounds.Update(delta)
	self.duration -= delta
}
