package debug

import (
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"
)

var debugFont font.Face

func init() {
	font, err := truetype.Parse(goregular.TTF)
	if err != nil {
		panic(err)
	}
	debugFont = truetype.NewFace(font, &truetype.Options{Size: 16})
}
