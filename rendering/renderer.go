package rendering

import (
	"github.com/SnareChops/nengine/types"
	"golang.org/x/exp/slices"
)

type orderedLayer struct {
	order int
	layer types.RenderLayer
}

// Renderer implements the Renderer interface
type Renderer struct {
	layers []orderedLayer
}

// Creates a new Renderer
func NewRenderer() *Renderer {
	return &Renderer{
		layers: []orderedLayer{},
	}
}

func (self *Renderer) AddRenderLayer(order int, layer types.RenderLayer) {
	// Avoid duplicates
	for _, l := range self.layers {
		if l.layer == layer {
			return
		}
	}
	self.layers = append(self.layers, orderedLayer{order, layer})
	slices.SortStableFunc(self.layers, func(a, b orderedLayer) int {
		return a.order - b.order
	})
}

func (self *Renderer) RemoveRenderLayer(layer types.RenderLayer) {
	for i, l := range self.layers {
		if l.layer == layer {
			self.layers = append(self.layers[:i], self.layers[i+1:]...)
			return
		}
	}
}

// Draw the result of the Renderer processing to the provided image
// (usually the screen)
func (self *Renderer) Draw(screen types.Image) {
	for _, l := range self.layers {
		l.layer.Draw(screen)
	}
}
