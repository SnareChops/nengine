package fonts

import (
	"os"
	"path"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

var (
	Arial   *truetype.Font
	Arial12 font.Face
	Arial14 font.Face
	Arial16 font.Face
	Arial18 font.Face
	Arial24 font.Face
	Arial28 font.Face
	Arial32 font.Face
	Arial36 font.Face
	Arial48 font.Face
)

func Init(root string) {
	arial, err := os.ReadFile(path.Join(root, "fonts/arial.ttf"))
	if err != nil {
		panic(err)
	}
	Arial, err = truetype.Parse(arial)
	if err != nil {
		panic(err)
	}

	Arial48 = truetype.NewFace(Arial, &truetype.Options{Size: 48})
	Arial36 = truetype.NewFace(Arial, &truetype.Options{Size: 36})
	Arial32 = truetype.NewFace(Arial, &truetype.Options{Size: 32})
	Arial28 = truetype.NewFace(Arial, &truetype.Options{Size: 28})
	Arial24 = truetype.NewFace(Arial, &truetype.Options{Size: 24})
	Arial18 = truetype.NewFace(Arial, &truetype.Options{Size: 18})
	Arial16 = truetype.NewFace(Arial, &truetype.Options{Size: 16})
	Arial14 = truetype.NewFace(Arial, &truetype.Options{Size: 14})
	Arial12 = truetype.NewFace(Arial, &truetype.Options{Size: 12})
}

func GetStringHeight(str string, face font.Face) (ascent int, descent int) {
	ascent = face.Metrics().Ascent.Ceil()
	descent = face.Metrics().Descent.Ceil()
	return
}

func GetStringWidth(str string, face font.Face) int {
	return font.MeasureString(face, str).Floor()
}

// GetStringSize returns the string measurements for the provided
// string using the provided font face.
// Returns (width, ascent, descent int)
func GetStringSize(str string, face font.Face) (int, int, int) {
	a, d := GetStringHeight(str, face)
	return GetStringWidth(str, face), a, d
}
