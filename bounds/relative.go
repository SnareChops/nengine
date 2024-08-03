package bounds

import (
	"github.com/SnareChops/nengine/types"
	"github.com/hajimehoshi/ebiten/v2"
)

// Relative represents a Bounds that is relative to another
// The position of this Bounds is added the "parent" Bounds to produce
// the final coordinates
type Relative struct {
	types.Box
	Parent  types.Box
	options *ebiten.DrawImageOptions
}

// Init the state of the RelativeBounds
func (self *Relative) Init(parent types.Box, width, height int) *Relative {
	self.Parent = parent
	self.Box = NewBox(width, height)
	self.options = &ebiten.DrawImageOptions{}
	return self
}

// Min returns the raw position of the top left corner of the bounds as (x, y float64)
func (self *Relative) Min() (float64, float64) {
	px, py := self.Parent.Pos2()
	x, y := self.Box.Min()
	return px + x, py + y
}

func (self *Relative) Mid() (float64, float64) {
	px, py := self.Parent.Pos2()
	x, y := self.Box.Mid()
	return px + x, py + y
}

func (self *Relative) Max() (float64, float64) {
	px, py := self.Parent.Pos2()
	x, y := self.Box.Max()
	return px + x, py + y
}

func (self *Relative) DrawOptions(camera types.Camera) *ebiten.DrawImageOptions {
	self.options.GeoM.Reset()
	rotation := self.Rotation()
	offx, offy := self.Offset()

	// Rotate around anchor
	if rotation != 0 {
		self.options.GeoM.Translate(-offx, -offy)
		self.options.GeoM.Rotate(rotation)
		self.options.GeoM.Translate(offx, offy)
	}
	// Translate
	if camera == nil {
		self.options.GeoM.Translate(self.Min())
		return self.options
	}
	x, y := camera.WorldToScreenPos(self.Min())
	self.options.GeoM.Translate(float64(x), float64(y))
	return self.options
}
