package shaders

import (
	_ "embed"

	"github.com/SnareChops/nengine/image"
	"github.com/SnareChops/nengine/types"
)

//go:embed fontShader.kage
var fontShader []byte
var FontShader types.Shader

func init() {
	var err error
	FontShader, err = image.NewShader(fontShader)
	if err != nil {
		panic(err)
	}
}
