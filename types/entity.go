package types

type Entity interface {
	Bounds
}

type UpdatableEntity interface {
	Entity
	Update(delta int)
}
