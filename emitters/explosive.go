package emitters

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/SnareChops/nengine/bounds"
	"github.com/SnareChops/nengine/debug"
	"github.com/SnareChops/nengine/types"
)

type Explosive struct {
	types.Bounds
	random               *rand.Rand
	particles            []Particle
	active               []Particle
	minVelocity          float64
	maxVelocity          float64
	minAngle             float64
	maxAngle             float64
	minLife              int
	maxLife              int
	density              int // How many particles to emit per 10ms
	duration             int // How long to emit for
	debugTimer           *debug.DebugTimer
	particleUpdateTimer  *debug.FrameTimer
	particleReleaseTimer *debug.FrameTimer
	particleRemoveTimer  *debug.FrameTimer
	removeTimer          *debug.FrameTimer
	despawnTimer         *debug.FrameTimer
}

func (self *Explosive) Init(particles []Particle) *Explosive {
	self.Bounds = new(bounds.Raw).Init(0, 0)
	self.random = rand.New(rand.NewSource(time.Now().UnixMilli()))
	self.particles = particles
	self.active = make([]Particle, len(self.particles))
	self.debugTimer = debug.NewDebugTimer("ExplosiveEmitter")
	self.particleUpdateTimer = debug.NewFrameTimer("Particle Update")
	self.particleReleaseTimer = debug.NewFrameTimer("Particle Release")
	self.particleRemoveTimer = debug.NewFrameTimer("Particle Remove")
	self.removeTimer = debug.NewFrameTimer("Remove")
	self.despawnTimer = debug.NewFrameTimer("Despawn")
	debug.DebugStat("Active Particles", func() string {
		count := 0
		for i := range self.active {
			if self.active[i] != nil {
				count++
			}
		}
		return fmt.Sprint(count)
	})
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
	self.removeTimer.Start()
	for i, p := range self.active {
		if p == particle {
			self.despawnTimer.Start()
			particle.Despawn()
			self.despawnTimer.End()
			self.active[i] = nil
			return
		}
	}
	self.removeTimer.End()
}

func (self *Explosive) Update(delta int) {
	self.debugTimer.Start()
	self.duration -= delta
	if self.duration <= 0 {
		self.duration = 0
		// return
	}
	desired := int((float64(delta) / 10) * float64(self.density))

	for _, particle := range self.particles {
		if particle.Duration() > 0 {
			self.particleUpdateTimer.Start()
			particle.Update(delta)
			if particle.Duration() <= 0 {
				self.particleRemoveTimer.Start()
				self.Remove(particle)
				self.particleRemoveTimer.End()
			}
			self.particleUpdateTimer.End()
		} else {
			if desired > 0 && self.duration > 0 {
				self.particleReleaseTimer.Start()
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
				self.particleReleaseTimer.End()
			}
		}
	}
	self.debugTimer.End()
	self.particleUpdateTimer.EndFrame()
	self.particleReleaseTimer.EndFrame()
	self.particleRemoveTimer.EndFrame()
	self.removeTimer.EndFrame()
	self.despawnTimer.EndFrame()
}
