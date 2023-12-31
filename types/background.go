package types

import "github.com/hajimehoshi/ebiten/v2"

type Background interface {
	ClearBackground()
	AddBackgroundImage(image *ebiten.Image, offsetX, offsetY float64)
}
