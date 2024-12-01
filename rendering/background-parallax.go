package rendering

import (
	"github.com/SnareChops/nengine/bounds"
	"github.com/SnareChops/nengine/types"
)

type ParallaxBackground struct {
	*bounds.Raw
	camera      types.Camera
	order       int
	worldWidth  int
	worldHeight int
	image       types.Image
}

func (self *ParallaxBackground) Init(order, viewWidth, viewHeight, worldWidth, worldHeight int, image types.Image) *ParallaxBackground {
	self.order = order
	self.worldWidth = worldWidth
	self.worldHeight = worldHeight
	self.image = image
	ww, wh := image.Size()
	self.Raw = new(bounds.Raw).Init(ww, wh)
	self.camera = new(BasicCamera).Init(viewWidth, viewHeight, ww, wh)
	return self
}

func (self *ParallaxBackground) Order() int {
	return self.order
}

func (self *ParallaxBackground) Update(x, y, delta int) {
	xp := float64(x) / float64(self.worldWidth)
	yp := float64(y) / float64(self.worldHeight)
	w, h := self.Size()
	self.camera.SetPos(float64(w)*xp, float64(h)*yp)
}

func (self *ParallaxBackground) Draw(screen types.Image) {
	screen.DrawImage(self.image.SubImage(self.camera.View()), nil)
}
