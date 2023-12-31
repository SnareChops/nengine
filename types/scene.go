package types

import "github.com/hajimehoshi/ebiten/v2"

type Scene interface {
	Update(delta int)
	Draw(screen *ebiten.Image)
}

type Initable interface {
	Init(Game)
}

type Loadable interface {
	Load(chan Scene, Game) Scene
}

type Destroyable interface {
	Destroy()
}
