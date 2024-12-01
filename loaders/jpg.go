package loaders

import (
	"image/jpeg"
	"os"

	"github.com/SnareChops/nengine/image"
	"github.com/SnareChops/nengine/types"
)

func LoadJPEG(path string) (types.Image, error) {
	background, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	img, err := jpeg.Decode(background)
	if err != nil {
		return nil, err
	}
	return image.NewImageFromImage(img), nil
}

func PreloadImageJpeg(alias, path string) {
	if _, ok := flat[alias]; ok {
		return
	}
	img, err := LoadJPEG(path)
	if err != nil {
		panic("PreloadImageJpeg: " + path + "\n" + err.Error())
	}
	flat[alias] = img
}

func PreloadSheetJpeg(alias, path string) {
	if _, ok := sheets[alias]; ok {
		return
	}
	width, height, err := detectSize(path)
	img, err := LoadJPEG(path)
	if err != nil {
		panic("PreloadSheetJpeg: " + path + "\n" + err.Error())
	}
	sheets[alias] = Sheet{
		CellWidth:  int(width),
		CellHeight: int(height),
		Cells:      slice(img, int(width), int(height)),
	}
}
