package fonts

import (
	"os"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

type ttfont struct {
	f *truetype.Font
	c map[float64]font.Face
}

func (f *ttfont) Face(size float64) font.Face {
	if face, ok := f.c[size]; ok {
		return face
	}
	f.c[size] = truetype.NewFace(f.f, &truetype.Options{Size: size})
	return f.c[size]
}

func LoadTTF(alias, path string) {
	if _, ok := fonts[alias]; ok {
		return
	}
	data, err := os.ReadFile(path)
	if err != nil {
		panic("LoadTTF: " + path + "\n" + err.Error())
	}
	f, err := truetype.Parse(data)
	if err != nil {
		panic("LoadTTF: " + path + "\n" + err.Error())
	}
	fonts[alias] = &ttfont{
		f: f,
		c: map[float64]font.Face{},
	}
}
