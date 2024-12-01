package ui

import (
	"image/color"
	"math"

	"github.com/SnareChops/nengine/bounds"
	"github.com/SnareChops/nengine/image"
	"github.com/SnareChops/nengine/types"
	"github.com/SnareChops/nengine/utils"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type PercentBar struct {
	*bounds.Raw

	image      types.Image
	value      float64
	borderSize int
	border     color.Color
	fill       color.Color
}

func (self *PercentBar) Init(w, h int, borderSize int, border, fill color.Color) *PercentBar {
	self.Raw = new(bounds.Raw).Init(w, h)
	self.image = image.NewImage(w, h)

	self.borderSize = borderSize

	self.border = border
	self.fill = fill

	self.render()

	return self
}

func (self *PercentBar) SetValue(value float64) {
	prev := self.value
	self.value = value

	if self.value != prev {
		self.render()
	}
}

func (self *PercentBar) render() {
	self.image.Clear()
	offset := float64(self.borderSize / 2)
	if int(offset)%2 != 0 || offset == 0 {
		offset += 1
	}

	w := utils.LinearInterpolate(0., float64(self.Dx()), self.value)
	vector.DrawFilledRect(self.image.Raw(), 0, 0, float32(w), float32(self.Dy()), self.fill, false)
	vector.StrokeRect(self.image.Raw(), float32(offset), float32(offset), float32(self.Dx())-float32(math.Ceil(offset*2)), float32(self.Dy())-float32(math.Ceil(offset*2)), float32(self.borderSize), self.border, false)

}

func (self *PercentBar) Image() types.Image {
	return self.image
}
