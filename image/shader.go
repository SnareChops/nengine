package image

import (
	"github.com/SnareChops/nengine/types"
	"github.com/hajimehoshi/ebiten/v2"
)

type shader ebiten.Shader

func NewShader(src []byte) (types.Shader, error) {
	return ebiten.NewShader(src)
}
