package font

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Font interface {
	Height() int
	Char(rune) *ebiten.Image
}

type Line struct {
	width   int
	height  int
	kerning int
	letters []Letter
}

type Letter struct {
	Image *ebiten.Image
	Char  rune
}
