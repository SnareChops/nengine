package animators

import "github.com/hajimehoshi/ebiten/v2"

type Animator interface {
	Start(name string)
	Update(delta int)
	Image() *ebiten.Image
}
