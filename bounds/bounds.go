package bounds

import "github.com/SnareChops/nengine/types"

// TODO: Might be an opportunity here to convert these to bitmask
const (
	TOP = iota
	CENTER
	BOTTOM
	LEFT
	RIGHT
)

type PhysicsBounds interface {
	types.Bounds
	Velocity() (float64, float64)
	SetRawVelocity(x, y float64)
	SetVelocity(angle, magnitude float64)
	Acceleration() (float64, float64)
	SetAcceleration(angle, magnitude float64)
	Update(delta int)
}
