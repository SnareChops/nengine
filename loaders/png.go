package loaders

import (
	"errors"
	"image"
	"image/png"
	"os"
	"regexp"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
)

func LoadPNG(path string) (*ebiten.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	img, err := png.Decode(file)
	if err != nil {
		return nil, err
	}
	return ebiten.NewImageFromImage(img), nil
}

func PreloadImagePng(alias, path string) {
	if _, ok := flat[alias]; ok {
		return
	}
	image, err := LoadPNG(path)
	if err != nil {
		panic("PreloadImagePng: " + path + "\n" + err.Error())
	}
	flat[alias] = ebiten.NewImageFromImage(image)
}

func PreloadSheetPng(alias, path string) {
	if _, ok := sheets[alias]; ok {
		return
	}
	width, height, err := detectSize(path)
	if err != nil {
		panic("PreloadSheetPng: " + path + "\n" + err.Error())
	}
	image, err := LoadPNG(path)
	if err != nil {
		panic("PreloadSheetPng: " + path + "\n" + err.Error())
	}
	sheets[alias] = Sheet{
		CellWidth:  int(width),
		CellHeight: int(height),
		Cells:      slice(image, int(width), int(height)),
	}
}

func detectSize(path string) (width, height int64, err error) {
	regex := regexp.MustCompile(`([a-zA-Z_/\\]+?)(\d+)x?(\d+)?`)
	matches := regex.FindStringSubmatch(path)
	if len(matches) == 0 {
		err = errors.New("Png sprite sheets must include the cell width and height in the filename. (ex: SpriteSheet32.png or SpriteSheet32x32.jpg)")
		return
	}
	width, err = strconv.ParseInt(matches[2], 10, 0)
	if err != nil {
		return
	}
	height = width
	if matches[3] != "" {
		height, err = strconv.ParseInt(matches[3], 10, 0)
		if err != nil {
			return
		}
	}
	return
}

func slice(img *ebiten.Image, cw, ch int) (images []*ebiten.Image) {
	cols := img.Bounds().Dx() / cw
	rows := img.Bounds().Dy() / ch
	for row := range rows {
		for col := range cols {
			x, y := col*cw, row*ch
			images = append(images, img.SubImage(image.Rect(x, y, x+cw, y+ch)).(*ebiten.Image))
		}
	}
	return
}
