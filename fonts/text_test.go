package fonts_test

import (
	"image/color"
	"testing"

	"github.com/SnareChops/nengine/bounds"
	"github.com/SnareChops/nengine/fonts"
	"github.com/golang/freetype/truetype"
	"github.com/stretchr/testify/assert"
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

func TestCenterAnchoringText(t *testing.T) {
	text := fonts.NewText("", fontFace, color.Black)

	w, h := text.Size()
	assert.Equal(t, 0, w)
	assert.Equal(t, 20, h)

	x, y := text.Offset()
	assert.Equal(t, 0., x)
	assert.Equal(t, 0., y)

	text.SetAnchor(bounds.CENTER, bounds.TOP)

	x, y = text.Offset()
	assert.Equal(t, 0., x)
	assert.Equal(t, 0., y)

	text.SetValue("test")
	x, y = text.Offset()
	assert.Equal(t, 13., x)
	assert.Equal(t, 0., y)

	x, y = text.Min()
	assert.Equal(t, -13., x)
	assert.Equal(t, 0., y)
}
