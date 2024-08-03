package types

import "github.com/hajimehoshi/ebiten/v2"

type Initable interface {
	Init(Game)
}

type Loadable interface {
	Load(chan Scene, Game) Scene
}

type Reloadable interface {
	Reload()
}

type Drawable interface {
	Draw(screen *ebiten.Image)
}

type Destroyable interface {
	Destroy()
}
