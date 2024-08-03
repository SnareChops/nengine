package types

import "github.com/hajimehoshi/ebiten/v2"

type Scene interface {
	Update(delta int)
	Draw(screen *ebiten.Image)
}
