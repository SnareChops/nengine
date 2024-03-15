package fonts

import "github.com/hajimehoshi/ebiten/v2"

type RasterFont struct {
	height int
	data   map[rune]*Char
}

type Char struct {
	Width int
	Image *ebiten.Image
}
