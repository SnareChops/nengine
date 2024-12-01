package types

type Animator interface {
	Start(name string)
	Update(delta int)
	Image() Image
}
