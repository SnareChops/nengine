package emitters

import (
	"math/rand"
	"time"

	"github.com/SnareChops/nengine/bounds"
	"github.com/SnareChops/nengine/types"
)

type Explosive struct {
	types.Bounds
	random      *rand.Rand
	particles   []Particle
	active      []Particle
	minVelocity float64
	maxVelocity float64
	minAngle    float64
	maxAngle    float64
	minLife     int
	maxLife     int
	density     int // How many particles to emit per 10ms
	duration    int // How long to emit for
}

func (self *Explosive) Init(particles []Particle) *Explosive {
	self.Bounds = new(bounds.Raw).Init(0, 0)
	self.random = rand.New(rand.NewSource(time.Now().UnixMilli()))
	self.particles = particles
	self.active = make([]Particle, len(self.particles))
	return self
}

func (self *Explosive) SetVelocity(min, max float64) {
	self.minVelocity = min
	self.maxVelocity = max
}

// SetAngle sets the minimum and maximum emittion angles in radians
func (self *Explosive) SetAngle(min, max float64) {
	self.minAngle = min
	self.maxAngle = max
}

func (self *Explosive) SetLife(min, max int) {
	if min <= 0 || max <= 0 {
		panic("Life must be larger than 0")
	}
	self.minLife = min
	self.maxLife = max
}

func (self *Explosive) SetDensity(density int) {
	self.density = density
}

func (self *Explosive) Particles() []Particle {
	return self.active
}

func (self *Explosive) StartEmitter(duration int) {
	self.duration = duration
}

func (self *Explosive) Remove(particle Particle) {
	for i, p := range self.active {
		if p == particle {
			particle.Despawn()
			self.active[i] = nil
			return
		}
	}
}

func (self *Explosive) Update(delta int) {
	self.duration -= delta
	if self.duration <= 0 {
		self.duration = 0
		// return
	}
	desired := int((float64(delta) / 10) * float64(self.density))

	for _, particle := range self.particles {
		if particle.Duration() > 0 {
			particle.Update(delta)
			if particle.Duration() <= 0 {
				self.Remove(particle)
			}
		} else {
			if desired > 0 && self.duration > 0 {
				particle.SetPos2(self.Pos2())
				particle.SetDuration(self.random.Intn(self.maxLife-self.minLife) + self.minLife)
				particle.SetVelocity(
					self.random.Float64()*(self.maxAngle-self.minAngle)+self.minAngle,
					self.random.Float64()*(self.maxVelocity-self.minVelocity)+self.minVelocity,
				)
				particle.Spawn()
				for i, active := range self.active {
					if active == nil {
						self.active[i] = particle
						break
					}
				}
				desired -= 1
			}
		}
	}
}
