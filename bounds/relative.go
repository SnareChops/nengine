package bounds

import "github.com/SnareChops/nengine/types"

// Relative represents a Bounds that is relative to another
// The position of this Bounds is added the "parent" Bounds to produce
// the final coordinates
type Relative struct {
	types.Bounds
	Parent types.Bounds
}

// Init the state of the RelativeBounds
func (self *Relative) Init(parent types.Bounds, width, height int) *Relative {
	self.Parent = parent
	self.Bounds = new(Raw).Init(width, height)
	return self
}

// RawPos returns the raw position of the top left corner of the bounds as (x, y float64)
func (self *Relative) RawPos() (float64, float64) {
	px, py := self.Parent.Pos2()
	x, y := self.Bounds.Pos2()
	ox, oy := self.Bounds.Offset()
	return px + x - ox, py + y - oy
}
