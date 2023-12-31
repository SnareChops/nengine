package emitters

import "github.com/SnareChops/nengine/types"

type Emitter interface {
	types.Bounds
	Particles() []Particle
	Update(delta int)
}
