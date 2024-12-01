package rendering

import (
	"github.com/SnareChops/nengine/image"
	"github.com/SnareChops/nengine/types"
)

type BufferedCamera struct {
	*BasicCamera
	options *types.DrawImageOptions
	image   types.Image
}

func (self *BufferedCamera) Init(viewWidth, viewHeight, worldWidth, worldHeight int) *BufferedCamera {
	self.BasicCamera = new(BasicCamera).Init(viewWidth, viewHeight, worldWidth, worldHeight)
	self.image = image.NewImage(viewWidth, viewHeight)
	return self
}

func (self *BufferedCamera) Image(source types.Image) types.Image {
	self.options.GeoM.Reset()
	self.options.GeoM.Scale(self.zoom, self.zoom)
	self.image.Clear()
	self.image.DrawImage(source.SubImage(self.View()), self.options)
	return self.image
}
