package types

import "github.com/hajimehoshi/ebiten/v2"

type Game interface {
	ebiten.Game
	LoadScene(scene Scene)
	Terminate()
}
