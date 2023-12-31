package rendering

import (
	"github.com/SnareChops/nengine/types"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/exp/slices"
)

// Renderer implements the Renderer interface
type Renderer struct {
	layers []types.RenderLayer
}

// Init sets the starting state for the Renderer
func (self *Renderer) Init() *Renderer {
	self.layers = []types.RenderLayer{}
	return self
}

func (self *Renderer) AddRenderLayer(layer types.RenderLayer) {
	if slices.Contains(self.layers, layer) {
		return
	}
	self.layers = append(self.layers, layer)
	slices.SortStableFunc(self.layers, func(a, b types.RenderLayer) int {
		return a.Order() - b.Order()
	})
}

func (self *Renderer) RemoveRenderLayer(layer types.RenderLayer) {
	for i, l := range self.layers {
		if l == layer {
			self.layers = append(self.layers[:i], self.layers[i+1:]...)
			return
		}
	}
}

// Draw the result of the Renderer processing to the provided image
// (usually the screen)
func (self *Renderer) Draw(screen *ebiten.Image) {
	for _, layer := range self.layers {
		layer.Draw(screen)
	}
}
