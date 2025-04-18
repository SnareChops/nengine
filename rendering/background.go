package rendering

import (
	"github.com/SnareChops/nengine/bounds"
	"github.com/SnareChops/nengine/types"
)

// Background represents an assortment of images to use for
// the background in the Renderer
// This is different than world and screen concepts because of image
// size limitations with the ebitengine library
// Consider using ChunkImage() or ChunkBounds() if needed to split a large image
// or area into smaller pieces
type Background struct {
	*bounds.Raw
	camera types.Camera
	order  int
	pieces []backgroundPiece
}

type backgroundPiece struct {
	image   types.Image
	options *types.DrawImageOptions
	x       float64
	y       float64
}

func (self *Background) Init(order, w, h int, camera types.Camera) *Background {
	self.order = order
	self.camera = camera
	self.Raw = new(bounds.Raw).Init(w, h)
	self.pieces = []backgroundPiece{}
	return self
}

func (self *Background) Order() int {
	return self.order
}

func (self *Background) ClearBackground() {
	self.pieces = []backgroundPiece{}
}

// AddBackgroundImage to the Background at the provided offset using world coordinates
func (self *Background) AddBackgroundImage(image types.Image, offsetX, offsetY float64) {
	self.pieces = append(self.pieces, backgroundPiece{
		image:   image,
		options: &types.DrawImageOptions{},
		x:       offsetX,
		y:       offsetY,
	})
}

func (self *Background) Draw(screen types.Image) {
	for _, piece := range self.pieces {
		if piece.image != nil {
			piece.options.GeoM.Reset()
			x, y := self.camera.WorldToScreenPos(piece.x, piece.y)
			piece.options.GeoM.Translate(float64(x), float64(y))
			screen.DrawImage(piece.image, piece.options)
		}
	}
}
