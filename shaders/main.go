package shaders

import (
	_ "embed"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed fontShader.kage
var fontShader []byte
var FontShader *ebiten.Shader

func init() {
	var err error
	FontShader, err = ebiten.NewShader(fontShader)
	if err != nil {
		panic(err)
	}
}
