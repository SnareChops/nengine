package types

type Background interface {
	ClearBackground()
	AddBackgroundImage(image Image, offsetX, offsetY float64)
}
