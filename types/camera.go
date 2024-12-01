package types

type Camera interface {
	Pos() (float64, float64)
	SetPos(x, y float64)
	ViewSize() (int, int)
	WorldSize() (int, int)
	View() (x, y, w, h int)
	Zoom() float64
	SetZoom(zoom float64)
	CursorWorldPosition() (float64, float64)
	WorldToScreenPos(x, y float64) (int, int)
	ScreenToWorldPos(screenX, screenY int) (float64, float64)
	Follow(target Position)
	Update(delta int)
}

type BufferedCamera interface {
	Image(source Image) Image
}
