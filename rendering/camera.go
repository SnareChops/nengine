package rendering

import (
	"fmt"
	"image"

	"github.com/SnareChops/nengine/types"
	"github.com/hajimehoshi/ebiten/v2"
)

// Camera represents a virtual camera for use by the Renderer
type Camera struct {
	x       int
	y       int
	vw      int
	vh      int
	ww      int
	wh      int
	zw      int
	zh      int
	rect    image.Rectangle
	zoom    float64
	options *ebiten.DrawImageOptions
	image   *ebiten.Image
	target  types.Entity
}

// Init the Camera with starting state
func (self *Camera) Init(viewWidth, viewHeight, worldWidth, worldHeight int) *Camera {
	self.zoom = 1
	self.vw = viewWidth
	self.vh = viewHeight
	self.ww = worldWidth
	self.wh = worldHeight
	self.zw = self.vw
	self.zh = self.vh
	self.options = &ebiten.DrawImageOptions{}
	self.image = ebiten.NewImage(viewWidth, viewHeight)
	self.SetCameraPos(0, 0)
	return self
}

func (self *Camera) CameraFollow(target types.Entity) {
	self.target = target
}

func (self *Camera) CameraPos() (int, int) {
	return self.x, self.y
}

// Set the position of the Camera
func (self *Camera) SetCameraPos(x, y int) {
	self.x, self.y = x, y
	self.resize()
}

func (self *Camera) resize() {
	// Calculate the points for the rectangle
	x1 := clamp(self.x-self.zw/2, 0, self.ww-self.zw)
	y1 := clamp(self.y-self.zh/2, 0, self.wh-self.zh)
	x2 := clamp(x1+self.zw, self.zw, self.ww)
	y2 := clamp(y1+self.zh, self.zh, self.wh)
	self.rect = image.Rect(x1, y1, x2, y2)
}

// ViewSize returns the view size of the Camera (width, height int)
func (self *Camera) CameraViewSize() (int, int) {
	return self.vw, self.vh
}

func (self *Camera) CameraWorldSize() (int, int) {
	return self.ww, self.wh
}

// View returns a rectangle representing the current view of the Camera
func (self *Camera) CameraView() image.Rectangle {
	return self.rect
}

func (self *Camera) CameraZoom() float64 {
	return self.zoom
}

func (self *Camera) SetCameraZoom(zoom float64) {
	if zoom <= 0 {
		fmt.Printf("Attempted to set camera zoom to an invalid number: %f\n", zoom)
		return
	}
	self.zoom = zoom
	self.zw = int(float64(self.vw) / self.zoom)
	self.zh = int(float64(self.vh) / self.zoom)
	self.resize()
}

func (self *Camera) CameraImage(source *ebiten.Image) *ebiten.Image {
	self.options.GeoM.Reset()
	self.options.GeoM.Scale(self.zoom, self.zoom)
	self.image.Clear()
	self.image.DrawImage(source.SubImage(self.rect).(*ebiten.Image), self.options)
	return self.image
}

func (self *Camera) CameraDraw(dest *ebiten.Image, source *ebiten.Image) {
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Scale(self.zoom, self.zoom)
	dest.DrawImage(source.SubImage(self.rect).(*ebiten.Image), options)
}

func (self *Camera) CursorWorldPosition() (float64, float64) {
	return self.ScreenToWorldPos(ebiten.CursorPosition())
}

// WorldToScreenPos converts the provided world coordinates to screen coordinates
// based on the current view of the Camera
func (self *Camera) WorldToScreenPos(x, y float64) (int, int) {
	return int(x) - self.rect.Min.X, int(y) - self.rect.Min.Y
}

// ScreenToWorldPos converts the provided screen coordinates to world coordinates
// based on the current view of the Camera
func (self *Camera) ScreenToWorldPos(screenX, screenY int) (float64, float64) {
	return float64(self.rect.Min.X) + float64(screenX)/self.zoom, float64(self.rect.Min.Y) + float64(screenY)/self.zoom
}

func (self *Camera) Update(delta int) {
	if self.target != nil {
		x, y := self.target.Pos2()
		self.SetCameraPos(int(x), int(y))
	}
}

func clamp(num, min, max int) int {
	if num < min {
		return min
	}
	if num > max {
		return max
	}
	return num
}
