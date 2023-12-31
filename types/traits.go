package types

import "github.com/hajimehoshi/ebiten/v2"

type Reloadable interface {
	Reload()
}

type Drawable interface {
	Draw(screen *ebiten.Image)
}
