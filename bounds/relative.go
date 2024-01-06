package bounds

import (
	"github.com/SnareChops/nengine/types"
	"github.com/hajimehoshi/ebiten/v2"
)

// Relative represents a Bounds that is relative to another
// The position of this Bounds is added the "parent" Bounds to produce
// the final coordinates
type Relative struct {
	types.Bounds
	Parent  types.Bounds
	options *ebiten.DrawImageOptions
}

// Init the state of the RelativeBounds
func (self *Relative) Init(parent types.Bounds, width, height int) *Relative {
	self.Parent = parent
	self.Bounds = new(Raw).Init(width, height)
	self.options = &ebiten.DrawImageOptions{}
	return self
}

// RawPos returns the raw position of the top left corner of the bounds as (x, y float64)
func (self *Relative) RawPos() (float64, float64) {
	px, py := self.Parent.Pos2()
	x, y := self.Bounds.Pos2()
	ox, oy := self.Bounds.Offset()
	return px + x - ox, py + y - oy
}

func (self *Relative) DrawOptions(camera types.Camera) *ebiten.DrawImageOptions {
	self.options.GeoM.Reset()
	rotation := self.Rotation()
	offx, offy := self.Offset()
	scalx, scaly := self.Scale()

	// Rotate around anchor
	if rotation != 0 {
		self.options.GeoM.Translate(-offx, -offy)
		self.options.GeoM.Rotate(rotation)
		self.options.GeoM.Translate(offx, offy)
	}
	// Scale
	if scalx != 1 && scaly != 1 {
		self.options.GeoM.Scale(scalx, scaly)
	}
	// Translate
	if camera == nil {
		self.options.GeoM.Translate(self.RawPos())
		return self.options
	}
	x, y := camera.WorldToScreenPos(self.RawPos())
	self.options.GeoM.Translate(float64(x), float64(y))
	return self.options
}
