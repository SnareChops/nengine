package rendering

import (
	"github.com/SnareChops/nengine/bounds"
	"github.com/hajimehoshi/ebiten/v2"
)

type ParallaxBackground struct {
	*bounds.Raw
	camera      *Camera
	order       int
	worldWidth  int
	worldHeight int
	image       *ebiten.Image
}

func (self *ParallaxBackground) Init(order, viewWidth, viewHeight, worldWidth, worldHeight int, image *ebiten.Image) *ParallaxBackground {
	self.order = order
	self.worldWidth = worldWidth
	self.worldHeight = worldHeight
	self.image = image
	ww, wh := image.Size()
	self.Raw = new(bounds.Raw).Init(ww, wh)
	self.camera = new(Camera).Init(viewWidth, viewHeight, ww, wh)
	return self
}

func (self *ParallaxBackground) Order() int {
	return self.order
}

func (self *ParallaxBackground) Update(x, y, delta int) {
	xp := float64(x) / float64(self.worldWidth)
	yp := float64(y) / float64(self.worldHeight)
	w, h := self.Size()
	x = int(float64(w) * xp)
	y = int(float64(h) * yp)
	self.camera.SetCameraPos(x, y)
}

func (self *ParallaxBackground) Draw(screen *ebiten.Image) {
	screen.DrawImage(self.image.SubImage(self.camera.CameraView()).(*ebiten.Image), nil)
}
