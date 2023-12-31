package types

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type Camera interface {
	CameraPos() (int, int)
	SetCameraPos(x, y int)
	CameraViewSize() (int, int)
	CameraView() image.Rectangle
	CameraZoom() float64
	SetCameraZoom(zoom float64)
	CameraImage(source *ebiten.Image) *ebiten.Image
	CursorWorldPosition() (float64, float64)
	WorldToScreenPos(x, y float64) (int, int)
	ScreenToWorldPos(screenX, screenY int) (float64, float64)
	CameraFollow(target Entity)
	Update(delta int)
}
