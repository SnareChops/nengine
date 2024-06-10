package fonts

import (
	"os"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

type otfont struct {
	f *opentype.Font
	c map[float64]font.Face
}

func (f *otfont) Face(size float64) font.Face {
	if face, ok := f.c[size]; ok {
		return face
	}
	var err error
	f.c[size], err = opentype.NewFace(f.f, &opentype.FaceOptions{Size: size})
	if err != nil {
		panic(err)
	}
	return f.c[size]
}

func LoadOTF(alias, path string) {
	if _, ok := fonts[alias]; ok {
		return
	}
	data, err := os.ReadFile(path)
	if err != nil {
		panic("LoadOTF: " + path + "\n" + err.Error())
	}
	f, err := opentype.Parse(data)
	if err != nil {
		panic("LoadOTF: " + path + "\n" + err.Error())
	}
	fonts[alias] = &otfont{
		f: f,
		c: map[float64]font.Face{},
	}
}
