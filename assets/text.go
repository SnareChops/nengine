package assets

import (
	"image/color"

	"github.com/SnareChops/nengine/bounds"
	"github.com/SnareChops/nengine/types"
	"github.com/hajimehoshi/ebiten/v2"
	ebitentext "github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

type VectorText struct {
	*bounds.Raw
	Value  string
	Face   font.Face
	Color  color.Color
	Ascent int
}

func (self *VectorText) Init(value string, face font.Face, color color.Color) *VectorText {
	self.Value = value
	self.Face = face
	self.Color = color
	self.Ascent = face.Metrics().Ascent.Ceil()
	w := font.MeasureString(face, value).Ceil()
	h := self.Ascent + face.Metrics().Descent.Ceil()
	self.Raw = new(bounds.Raw).Init(w, h)
	return self
}

func DrawVectorText(screen *ebiten.Image, text *VectorText, camera types.Camera) {
	x, y := int(text.X()), int(text.Y())
	if camera != nil {
		x, y = camera.WorldToScreenPos(text.Pos2())
	}
	y += text.Ascent
	ebitentext.Draw(screen, text.Value, text.Face, x, y, text.Color)
}
