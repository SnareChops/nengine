package types

type Entity interface {
	Bounds
	Update(delta int)
}
