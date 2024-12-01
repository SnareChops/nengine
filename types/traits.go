package types

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
	Draw(screen Image)
}

type Destroyable interface {
	Destroy()
}
