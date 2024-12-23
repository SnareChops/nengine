package bounds

import (
	"github.com/SnareChops/nengine/types"
	"github.com/hajimehoshi/ebiten/v2"
)

// Raw A Bounds instance that represents an absolutely
// positioned bounds
type Raw struct {
	*Box
	options *types.DrawImageOptions
}

// Init sets the initial state of the RawBounds
func (self *Raw) Init(width, height int) *Raw {
	self.Box = NewBox(width, height)
	self.options = &ebiten.DrawImageOptions{}
	return self
}

func (self *Raw) InitFromPoints(a types.Position, b types.Position) *Raw {
	x1 := min(a.X(), b.X())
	y1 := min(a.Y(), b.Y())
	x2 := max(a.X(), b.X())
	y2 := max(a.Y(), b.Y())
	self.Init(int(x2-x1), int(y2-y1))
	self.SetPos2(x1, y1)
	return self
}

func (self *Raw) DrawOptions(sx, sy float64, camera types.Camera) *types.DrawImageOptions {
	self.options.GeoM.Reset()
	rotation := self.Rotation()
	offx, offy := self.Offset()

	// Rotate around anchor
	if rotation != 0 {
		self.options.GeoM.Translate(-offx, -offy)
		self.options.GeoM.Rotate(rotation)
		self.options.GeoM.Translate(offx, offy)
	}
	x, y := self.Min()
	if self.fx < 0 {
		x += float64(self.Dx())
	}
	if self.fy < 0 {
		y += float64(self.Dy())
	}
	// If camera is provided, scale and translate
	if camera != nil {
		self.options.GeoM.Scale(sx*self.fx*camera.Zoom(), sy*self.fy*camera.Zoom())
		x, y := camera.WorldToScreenPos(x, y)
		self.options.GeoM.Translate(float64(x), float64(y))
		return self.options
	}
	// Translate
	self.options.GeoM.Scale(sx*self.fx, sy*self.fy)
	self.options.GeoM.Translate(x, y)
	return self.options
}
