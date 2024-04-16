package cache

import (
	"fmt"
	"image"
	"regexp"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
)

type sheet struct {
	image image.Image
	cw    int
	ch    int
}

var sheets = map[string]sheet{}

func IsSheet(filename string) (bool, string, int, int) {
	regex := regexp.MustCompile(`([a-zA-Z_/\\]+?)(\d+)x?(\d+)?`)
	matches := regex.FindStringSubmatch(filename)
	if len(matches) == 0 {
		return false, "", 0, 0
	}
	width, err := strconv.ParseInt(matches[2], 10, 0)
	if err != nil {
		fmt.Printf("Failed to parse sheet cell width: %s %s %s\n", filename, matches[2], err.Error())
		return false, "", 0, 0
	}
	var height int64
	if matches[3] != "" {
		height, err = strconv.ParseInt(matches[3], 10, 0)
		if err != nil {
			fmt.Printf("Failed to parse sheet cell height: %s %s %s\n", filename, matches[3], err.Error())
			return false, "", 0, 0
		}
	} else {
		height = width
	}

	return true, matches[0], int(width), int(height)
}

func addSheet(alias string, width, height int, image image.Image) {
	sheets[alias] = sheet{
		image: image,
		cw:    width,
		ch:    height,
	}
}

type SubImageable interface {
	SubImage(r image.Rectangle) image.Image
}

func PromoteSheet(alias string) (int, int, int, int, []*ebiten.Image) {
	sheet := sheets[alias]
	if sheet.image == nil {
		panic("Sheet is not in cache: " + alias)
	}
	images := []*ebiten.Image{}
	cols := sheet.image.Bounds().Dx() / sheet.cw
	rows := sheet.image.Bounds().Dy() / sheet.ch
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			x, y := col*sheet.cw, row*sheet.ch
			img := ebiten.NewImageFromImage(sheet.image.(*image.NRGBA).SubImage(image.Rect(x, y, x+sheet.cw, y+sheet.ch)))
			images = append(images, img)
		}
	}
	delete(sheets, alias)
	return sheet.image.Bounds().Dx(), sheet.image.Bounds().Dy(), sheet.cw, sheet.ch, images
}
