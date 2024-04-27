package font

// DON"T USE THIS FILE, USE assets/text.go
import (
	"image/color"
	"strings"

	"github.com/SnareChops/nengine/bounds"
	"github.com/hajimehoshi/ebiten/v2"
	ebitentext "github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

type Text struct {
	*bounds.Raw
	metrics font.Metrics
	face    font.Face
	lines   []string
	color   color.Color
}

func (self *Text) resize() {
	metrics := self.face.Metrics()
	height := metrics.Ascent.Ceil() + metrics.Descent.Ceil()
	max := 0
	for _, line := range self.lines {
		w := font.MeasureString(self.face, line).Ceil()
		if w > max {
			max = w
		}
	}
	self.Raw.Resize(max, height*len(self.lines))
}

func NewText(text string, face font.Face) *Text {
	result := &Text{
		Raw:     new(bounds.Raw).Init(0, 0),
		metrics: face.Metrics(),
		face:    face,
		lines:   strings.Split(text, "\n"),
	}
	result.resize()
	return result
}

func DrawVectorText(dest *ebiten.Image, text *Text) {
	for i, line := range text.lines {
		ascent := text.metrics.Ascent.Ceil()
		height := ascent + text.metrics.Descent.Ceil()
		ebitentext.Draw(dest, line, text.face, int(text.X()), int(text.Y())+i*height+ascent, text.color)
	}
	ebitentext.Draw(dest, text.lines[0], text.face, int(text.X()), int(text.Y()), text.color)
}

func Wrap(width int, text *Text) {
	var curr string
	var wrapped []string
	for _, line := range text.lines {
		words := strings.Split(line, " ")
		for i := 0; i < len(words); i++ {
			if i == 0 {
				curr = words[i]
			} else {
				curr += " " + words[i]
			}
			if font.MeasureString(text.face, curr).Ceil() > width {
				wrapped = append(wrapped, strings.Join(words[:i], " "))
				words = words[i:]
				curr = ""
				i = -1
			}
		}
		wrapped = append(wrapped, strings.Join(words, " "))
	}
	text.lines = wrapped
	text.resize()
}
