package fonts

import (
	"github.com/SnareChops/nengine/types"
	"github.com/hajimehoshi/ebiten/v2"
	ebitentext "github.com/hajimehoshi/ebiten/v2/text"
)

func DrawText(screen *ebiten.Image, text *Text, camera types.Camera) {
	x, y := int(text.MinX()), int(text.MinY())
	if camera != nil {
		x, y = camera.WorldToScreenPos(text.Pos2())
	}
	y += text.Ascent.Ceil()
	height := text.Descent.Ceil() + text.Ascent.Ceil()
	for i, line := range text.Lines {
		ebitentext.Draw(screen, line, text.Face, x, y+i*height, text.Color)
	}
}
