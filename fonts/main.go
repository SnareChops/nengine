package fonts

import (
	"strings"

	"golang.org/x/image/font"
)

type ft interface {
	Face(size float64) font.Face
}

var fonts = map[string]ft{}

func LoadFont(alias, path string) {
	p := strings.ToLower(path)
	if strings.HasSuffix(p, ".ttf") {
		LoadTTF(alias, path)
		return
	}
	if strings.HasSuffix(p, ".otf") {
		LoadOTF(alias, path)
		return
	}
	panic("LoadFont: Unsupported font type")
}

func Font(alias string, size float64) font.Face {
	if font, ok := fonts[alias]; ok {
		return font.Face(size)
	}
	panic("Font: " + alias + " not loaded")
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
