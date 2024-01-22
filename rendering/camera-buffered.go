package rendering

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type BufferedCamera struct {
	*BasicCamera
	options *ebiten.DrawImageOptions
	image   *ebiten.Image
}

func (self *BufferedCamera) Init(viewWidth, viewHeight, worldWidth, worldHeight int) *BufferedCamera {
	self.BasicCamera = new(BasicCamera).Init(viewWidth, viewHeight, worldWidth, worldHeight)
	return self
}

func (self *BufferedCamera) Image(source *ebiten.Image) *ebiten.Image {
	self.options.GeoM.Reset()
	self.options.GeoM.Scale(self.zoom, self.zoom)
	self.image.Clear()
	self.image.DrawImage(source.SubImage(self.View()).(*ebiten.Image), self.options)
	return self.image
}
