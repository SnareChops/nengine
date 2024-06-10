package loaders

import (
	"image/jpeg"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

func LoadJPEG(path string) (*ebiten.Image, error) {
	background, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	img, err := jpeg.Decode(background)
	if err != nil {
		return nil, err
	}
	return ebiten.NewImageFromImage(img), nil
}

func PreloadImageJpeg(alias, path string) {
	if _, ok := flat[alias]; ok {
		return
	}
	image, err := LoadJPEG(path)
	if err != nil {
		panic("PreloadImageJpeg: " + path + "\n" + err.Error())
	}
	flat[alias] = ebiten.NewImageFromImage(image)
}

func PreloadSheetJpeg(alias, path string) {
	if _, ok := sheets[alias]; ok {
		return
	}
	width, height, err := detectSize(path)
	image, err := LoadJPEG(path)
	if err != nil {
		panic("PreloadSheetJpeg: " + path + "\n" + err.Error())
	}
	sheets[alias] = Sheet{
		CellWidth:  int(width),
		CellHeight: int(height),
		Cells:      slice(image, int(width), int(height)),
	}
}
