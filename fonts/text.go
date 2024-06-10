package fonts

import (
	"image/color"
	"strings"

	"github.com/SnareChops/nengine/bounds"
	"golang.org/x/image/font"
)

type Text struct {
	*bounds.Raw
	font.Metrics
	Lines []string
	Face  font.Face
	Color color.Color
}

func NewText(value string, face font.Face, color color.Color) *Text {
	t := &Text{
		Lines:   strings.Split(value, "\n"),
		Face:    face,
		Color:   color,
		Metrics: face.Metrics(),
		Raw:     new(bounds.Raw).Init(0, 0),
	}
	t.resize()
	return t
}

func (self *Text) Wrap(width int) {
	var curr string
	var wrapped []string
	for _, line := range self.Lines {
		words := strings.Split(line, " ")
		for i := 0; i < len(words); i++ {
			if i == 0 {
				curr = words[i]
			} else {
				curr += " " + words[i]
			}
			if font.MeasureString(self.Face, curr).Ceil() > width {
				wrapped = append(wrapped, strings.Join(words[:i], " "))
				words = words[i:]
				curr = ""
				i = -1
			}
		}
		wrapped = append(wrapped, strings.Join(words, " "))
	}
	self.Lines = wrapped
	self.resize()
}

func (self *Text) resize() {
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
