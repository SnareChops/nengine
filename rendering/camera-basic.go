package rendering

import (
	"image"

	"github.com/SnareChops/nengine/types"
	"github.com/hajimehoshi/ebiten/v2"
)

type BasicCamera struct {
	*CameraBounds
	ww, wh int
	zoom   float64
	target types.Bounds
}

func (self *BasicCamera) Init(viewWidth, viewHeight, worldWidth, worldHeight int) *BasicCamera {
	self.zoom = 1
	self.ww = worldWidth
	self.wh = worldHeight
	self.CameraBounds = new(CameraBounds).Init(viewWidth, viewHeight)
	self.SetPos(0, 0)
	return self
}

func (self *BasicCamera) SetPos(x, y float64) {
	self.CameraBounds.SetPos(x, y)
	if ax := self.MinX(); ax < 0 {
		self.CameraBounds.SetPos(x+(0-ax), self.Y())
	}
	if ay := self.MinY(); ay < 0 {
		self.CameraBounds.SetPos(self.X(), y+(0-ay))
	}
	ww, wh := float64(self.ww), float64(self.wh)
	if ax := self.MaxX(); ax > ww {
		self.CameraBounds.SetPos(x-(ax-ww), self.Y())
	}
	if ay := self.MaxY(); ay > wh {
		self.CameraBounds.SetPos(self.X(), y-(ay-wh))
	}
}

func (self *BasicCamera) Zoom() float64 {
	return self.zoom
}

func (self *BasicCamera) SetZoom(zoom float64) {
	self.zoom = zoom
	self.Resize(int(float64(self.Dx())/zoom), int(float64(self.Dy())/zoom))
}

func (self *BasicCamera) Follow(target types.Bounds) {
	self.target = target
}

// ViewSize returns the view size of the Camera (width, height int)
func (self *BasicCamera) ViewSize() (int, int) {
	return self.Size()
}

func (self *BasicCamera) WorldSize() (int, int) {
	return self.ww, self.wh
}

func (self *BasicCamera) View() image.Rectangle {
	x1, y1 := self.Min()
	x2, y2 := self.Max()
	return image.Rect(int(x1), int(y1), int(x2), int(y2))
}

func (self *BasicCamera) CursorWorldPosition() (float64, float64) {
	return self.ScreenToWorldPos(ebiten.CursorPosition())
}

// WorldToScreenPos converts the provided world coordinates to screen coordinates
// based on the current view of the Camera
func (self *BasicCamera) WorldToScreenPos(x, y float64) (int, int) {
	return int((x - self.MinX()) * self.zoom), int((y - self.MinY()) * self.zoom)
}

// ScreenToWorldPos converts the provided screen coordinates to world coordinates
// based on the current view of the Camera
func (self *BasicCamera) ScreenToWorldPos(screenX, screenY int) (float64, float64) {
	return self.MinX() + float64(screenX)/self.zoom, self.MinY() + float64(screenY)/self.zoom
}

func (self *BasicCamera) Update(delta int) {
	if self.target != nil {
		self.SetPos(self.target.Pos2())
	}
}
