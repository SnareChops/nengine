package assets

import (
	"image/color"
	"strings"

	"github.com/SnareChops/nengine/bounds"
	"github.com/SnareChops/nengine/types"
	"github.com/hajimehoshi/ebiten/v2"
	ebitentext "github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

type VectorText struct {
	*bounds.Raw
	font.Metrics
	Lines []string
	Face  font.Face
	Color color.Color
}

func (self *VectorText) Init(value string, face font.Face, color color.Color) *VectorText {
	self.Lines = strings.Split(value, "\n")
	self.Face = face
	self.Color = color
	self.Metrics = face.Metrics()
	self.Raw = new(bounds.Raw).Init(0, 0)
	self.resize()
	return self
}

func (self *VectorText) resize() {
	height := self.Ascent.Ceil() + self.Descent.Ceil()
	max := 0
	for _, line := range self.Lines {
		w := font.MeasureString(self.Face, line).Ceil()
		if w > max {
			max = w
		}
	}
	self.Raw.Resize(max, height*len(self.Lines))
}

func DrawVectorText(screen *ebiten.Image, text *VectorText, camera types.Camera) {
	x, y := int(text.X()), int(text.Y())
	if camera != nil {
		x, y = camera.WorldToScreenPos(text.Pos2())
	}
	y += text.Ascent.Ceil()
	height := text.Descent.Ceil() + text.Ascent.Ceil()
	for i, line := range text.Lines {
		ebitentext.Draw(screen, line, text.Face, x, y+i*height, text.Color)
	}
}

func VectorTextWrap(width int, text *VectorText) {
	var curr string
	var wrapped []string
	for _, line := range text.Lines {
		words := strings.Split(line, " ")
		for i := 0; i < len(words); i++ {
			if i == 0 {
				curr = words[i]
			} else {
				curr += " " + words[i]
			}
			if font.MeasureString(text.Face, curr).Ceil() > width {
				wrapped = append(wrapped, strings.Join(words[:i], " "))
				words = words[i:]
				curr = ""
				i = -1
			}
		}
		wrapped = append(wrapped, strings.Join(words, " "))
	}
	text.Lines = wrapped
	text.resize()
}
