package rendering

import (
	"fmt"

	"github.com/SnareChops/nengine/bounds"
	"github.com/SnareChops/nengine/debug"
	"github.com/SnareChops/nengine/types"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/exp/slices"
)

type Screen struct {
	*bounds.Raw
	order   int
	sprites []types.Sprite
}

func (self *Screen) Init(order, width, height int) *Screen {
	self.Raw = new(bounds.Raw).Init(width, height)
	self.order = order
	self.sprites = []types.Sprite{}
	debug.DebugStat("Screen", func() string {
		x, y := ebiten.CursorPosition()
		return fmt.Sprintf("%d, %d", x, y)
	})
	return self
}

func (self *Screen) Order() int {
	return self.order
}

func (self *Screen) Sprites() []types.Sprite {
	return self.sprites
}

func (self *Screen) AddSprite(sprite types.Sprite) {
	if slices.Contains(self.sprites, sprite) {
		return
	}
	self.sprites = append(self.sprites, sprite)
}

func (self *Screen) RemoveSprite(sprite types.Sprite) {
	for i, s := range self.sprites {
		if s == sprite {
			self.sprites = append(self.sprites[:i], self.sprites[i+1:]...)
			return
		}
	}
}

func (self *Screen) Draw(screen types.Image) {
	slices.SortStableFunc(self.sprites, func(a, b types.Sprite) int {
		_, _, az := a.Pos3()
		_, _, bz := b.Pos3()
		return int(az - bz)
	})

	for _, sprite := range self.sprites {
		DrawSprite(screen, sprite, nil)
	}
}
