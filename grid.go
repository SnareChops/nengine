package nengine

import (
	"image/color"

	"github.com/SnareChops/nengine/image"
	"github.com/SnareChops/nengine/loaders"
	"github.com/SnareChops/nengine/rendering"
	"github.com/SnareChops/nengine/types"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type SpriteGrid struct {
	*RawBounds
	cw            int
	ch            int
	contents      []types.Sprite
	lines         []Line
	lineColor     *color.Color
	lineWidth     float32
	selectedColor color.Color
	selectedWidth float32
	selected      int
	image         types.Image
}

func (self *SpriteGrid) Init(gridWidth, gridHeight, cellWidth, cellHeight int) *SpriteGrid {
	width := gridWidth / cellWidth * cellWidth
	height := gridHeight / cellHeight * cellHeight
	self.cw = cellWidth
	self.ch = cellHeight
	self.selected = -1
	self.RawBounds = new(RawBounds).Init(width, height)
	self.contents = make([]types.Sprite, (width/cellWidth)*(height/cellHeight))
	self.lines = []Line{}
	for x := 0; x < width; x += cellWidth {
		self.lines = append(self.lines, NewLine(x, 0, x, height))
	}
	for y := 0; y < height; y += cellHeight {
		self.lines = append(self.lines, NewLine(0, y, width, y))
	}
	self.image = image.NewImage(self.Size())
	return self
}

func (self *SpriteGrid) InitFromTileSheet(sheet loaders.Sheet) *SpriteGrid {
	self.Init(sheet.Width, sheet.Height, sheet.CellWidth, sheet.CellHeight)
	for i, image := range sheet.Cells {
		self.SetContent(i, new(SimpleSprite).Init(image))
	}
	return self
}

func (self *SpriteGrid) Len() int {
	return len(self.contents)
}

func (self *SpriteGrid) Resize(w, h int) {
	wf := float64(w) / float64(self.Dx())
	hf := float64(h) / float64(self.Dy())
	scale := min(wf, hf)
	self.cw = int(float64(self.cw) * scale)
	self.ch = int(float64(self.ch) * scale)
	self.RawBounds.Resize(w, h)
	for _, content := range self.contents {
		if content != nil {
			content.Resize(self.cw, self.ch)
		}
	}
	self.image = image.NewImage(self.Size())
	self.render()
}

func (self *SpriteGrid) CellSize() (int, int) {
	return self.cw, self.ch
}

func (self *SpriteGrid) Contents() []types.Sprite {
	return self.contents
}

func (self *SpriteGrid) GetContent(index int) types.Sprite {
	if index < 0 || index >= len(self.contents) {
		return nil
	}
	return self.contents[index]
}

func (self *SpriteGrid) SetContent(index int, content types.Sprite) {
	self.contents[index] = content
	if content != nil {
		x, y := Floats(self.IndexPos(index))
		content.SetAnchor(LEFT, TOP)
		content.SetPos2(x, y)
		content.Resize(self.cw, self.ch)
	}
	self.render()
}

func (self *SpriteGrid) SetAllContent(contents []types.Sprite) {
	self.contents = contents
	for i, content := range self.contents {
		if content != nil {
			x, y := Floats(self.IndexPos(i))
			content.SetAnchor(LEFT, TOP)
			content.SetPos2(x, y)
			content.Resize(self.cw, self.ch)
		}
	}
	self.render()
}

func (self *SpriteGrid) SetRawContent(index int, image types.Image) {
	self.SetContent(index, new(SimpleSprite).Init(image))
}

func (self *SpriteGrid) IndexAt(x, y int) int {
	return (y/self.ch)*(self.Dx()/self.cw) + (x / self.cw)
}

func (self *SpriteGrid) IndexPos(index int) (int, int) {
	x := index * self.cw % self.Dx()
	y := int(float64(index)*float64(self.cw)/float64(self.Dx())) * self.ch
	return x, y
}

func (self *SpriteGrid) GetContentAt(x, y int) types.Sprite {
	return self.GetContent(self.IndexAt(x, y))
}

func (self *SpriteGrid) SetContentAt(x, y int, content types.Sprite) {
	self.SetContent(self.IndexAt(x, y), content)
}

func (self *SpriteGrid) SetRawContentAt(x, y int, image types.Image) {
	self.SetContent(self.IndexAt(x, y), new(SimpleSprite).Init(image))
}

func (self *SpriteGrid) Select(index int, width float32, color color.Color) {
	self.selected = index
	self.selectedWidth = width
	self.selectedColor = color
	self.render()
}

func (self *SpriteGrid) SelectAt(x, y int, width float32, color color.Color) {
	self.Select(self.IndexAt(x, y), width, color)
}

func (self *SpriteGrid) Reload() {
	for _, sprite := range self.contents {
		if sprite, ok := sprite.(types.Reloadable); ok {
			sprite.Reload()
		}
	}
	self.render()
}

func (self *SpriteGrid) Lines() []Line {
	return self.lines
}

func (self *SpriteGrid) ShowLines(width float32, color color.Color) {
	self.lineColor = &color
	self.lineWidth = width
	self.render()
}

func (self *SpriteGrid) HideLines() {
	self.lineColor = nil
	self.render()
}

func (self *SpriteGrid) ToggleLines(width float32, color color.Color) {
	if self.lineColor != nil {
		self.HideLines()
	} else {
		self.ShowLines(width, color)
	}
}

func (self *SpriteGrid) render() {
	self.image.Clear()
	for _, content := range self.contents {
		if content != nil {
			rendering.DrawSprite(self.image, content, nil)
		}
	}
	if self.lineColor != nil {
		for _, line := range self.lines {
			StrokeLine(self.image, line, self.lineWidth, *self.lineColor, false, nil)
		}
	}
	if self.selected >= 0 {
		x, y := self.IndexPos(self.selected)
		vector.StrokeRect(self.image.Raw(), float32(x), float32(y), float32(self.cw), float32(self.ch), self.selectedWidth, self.selectedColor, false)
	}
}

func (self *SpriteGrid) Image() types.Image {
	return self.image
}
