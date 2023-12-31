package nengine

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type SpriteSource interface {
	Alias() string
	Index() int
	Image() *ebiten.Image
}

type SpriteSourceGrid struct {
	*SpriteGrid
	sources []SpriteSource
}

func (self *SpriteSourceGrid) Init(gridWidth, gridHeight, cellWidth, cellHeight int) *SpriteSourceGrid {
	self.SpriteGrid = new(SpriteGrid).Init(gridWidth, gridHeight, cellWidth, cellHeight)
	self.sources = make([]SpriteSource, len(self.contents))
	return self
}

func (self *SpriteSourceGrid) Sources() []SpriteSource {
	return self.sources
}

func (self *SpriteSourceGrid) GetSource(index int) SpriteSource {
	return self.sources[index]
}

func (self *SpriteSourceGrid) GetSourceAt(x, y int) SpriteSource {
	return self.sources[self.IndexAt(x, y)]
}

func (self *SpriteSourceGrid) SetContent(index int, source SpriteSource) {
	self.sources[index] = source
	if source != nil {
		self.SpriteGrid.SetContent(index, new(SourceSprite).Init(source))
	} else {
		self.SpriteGrid.SetContent(index, nil)
	}
}

func (self *SpriteSourceGrid) SetAllContent(contents []SpriteSource) {
	sprites := []Sprite{}
	for _, source := range contents {
		if source != nil {
			sprites = append(sprites, new(SourceSprite).Init(source))
		} else {
			sprites = append(sprites, nil)
		}
	}
	self.SpriteGrid.SetAllContent(sprites)
}

func (self *SpriteSourceGrid) SetContentAt(x, y int, source SpriteSource) {
	self.SetContent(self.IndexAt(x, y), source)
}
