package types

type Input interface {
	RenderLayer
	InputCapture()
	IsInputCaptured() bool
	CursorContent() Sprite
	SetCursorContent(content Sprite)
	CursorDelta() (int, int)
	Update()
}
