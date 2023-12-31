package cache

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

var images = map[string]image.Image{}

func addImage(alias string, image image.Image) {
	images[alias] = image
}

func PromoteImage(alias string) *ebiten.Image {
	image := images[alias]
	delete(images, alias)
	return ebiten.NewImageFromImage(image)
}
