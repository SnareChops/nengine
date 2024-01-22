package types

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type Camera interface {
	Pos() (float64, float64)
	SetPos(x, y float64)
	ViewSize() (int, int)
	WorldSize() (int, int)
	View() image.Rectangle
	Zoom() float64
	SetZoom(zoom float64)
	CursorWorldPosition() (float64, float64)
	WorldToScreenPos(x, y float64) (int, int)
	ScreenToWorldPos(screenX, screenY int) (float64, float64)
	Follow(target Bounds)
	Update(delta int)
}

type BufferedCamera interface {
	Image(source *ebiten.Image) *ebiten.Image
}
