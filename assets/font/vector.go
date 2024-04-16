package font

import (
	"image/color"

	"github.com/SnareChops/nengine/bounds"
	"github.com/hajimehoshi/ebiten/v2"
	ebitentext "github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

type Text struct {
	*bounds.Raw
	face  font.Face
	text  string
	color color.Color
}

func NewText(text string, face font.Face) *Text {
	metrics := face.Metrics()
	height := metrics.Ascent.Ceil() + metrics.Descent.Ceil()
	width := font.MeasureString(face, text).Ceil()
	return &Text{
		Raw:  new(bounds.Raw).Init(width, height),
		face: face,
		text: text,
	}
}

func DrawVectorText(dest *ebiten.Image, text *Text) {
	ebitentext.Draw(dest, text.text, text.face, int(text.X()), int(text.Y()), text.color)
}
