package rendering

import (
	"fmt"
	"slices"

	"github.com/SnareChops/nengine/bounds"
	"github.com/SnareChops/nengine/debug"
	"github.com/SnareChops/nengine/types"
)

type World struct {
	*bounds.Raw
	order   int
	camera  types.Camera
	sprites []types.Sprite
}

func (self *World) Init(order int, camera types.Camera) *World {
	self.Raw = new(bounds.Raw).Init(camera.WorldSize())
	self.order = order
	self.camera = camera
	self.sprites = []types.Sprite{}
	debug.DebugStat("World", func() string {
		x, y := self.camera.CursorWorldPosition()
		return fmt.Sprintf("%.2f, %.2f", x, y)
	})
	return self
}

func (self *World) Order() int {
	return self.order
}

func (self *World) Sprites() []types.Sprite {
	return self.sprites
}

func (self *World) AddSprite(sprite types.Sprite) {
	if slices.Contains(self.sprites, sprite) {
		return
	}
	self.sprites = append(self.sprites, sprite)
}

func (self *World) RemoveSprite(sprite types.Sprite) {
	for i, s := range self.sprites {
		if s == sprite {
			self.sprites = append(self.sprites[:i], self.sprites[i+1:]...)
			return
		}
	}
}

func (self *World) Draw(screen types.Image) {
	slices.SortStableFunc(self.sprites, func(a, b types.Sprite) int {
		_, ay := a.Pos2()
		_, by := b.Pos2()
		return int(ay - by)
	})
	for _, sprite := range self.sprites {
		DrawSprite(screen, sprite, self.camera)
	}
}
