package types

type Scene interface {
	Update(delta int)
	Draw(screen Image)
}
