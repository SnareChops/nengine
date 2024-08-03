package emitters

import (
	"github.com/SnareChops/nengine/bounds"
	"github.com/SnareChops/nengine/types"
)

type Particle interface {
	types.PhysicsBounds
	Update(delta int)
	Duration() int
	SetDuration(ms int)
	Spawn()
	Despawn()
}

type ParticleBase struct {
	types.PhysicsBounds
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
