package console

import (
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"
)

var fontFace font.Face

func init() {
	f, err := truetype.Parse(goregular.TTF)
	if err != nil {
		panic(err)
	}
	fontFace = truetype.NewFace(f, &truetype.Options{Size: 16})
}
