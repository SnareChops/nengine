package assets

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/SnareChops/nengine/assets/cache"
	"github.com/hajimehoshi/ebiten/v2"
)

type ImageSource struct {
	alias string
}

func NewImageSource(alias string) ImageSource {
	return ImageSource{alias}
}

func (self ImageSource) Alias() string {
	return self.alias
}

func (self ImageSource) Index() int {
	return 0
}

func (self ImageSource) Image() *ebiten.Image {
	return GetImage(self.alias)
}

var images = map[string]*ebiten.Image{}

func GetImage(alias string) *ebiten.Image {
	if image, ok := images[alias]; ok {
		return image
	}
	image := cache.PromoteImage(alias)
	images[alias] = image
	return image
}

func LoadImage(path string) *ebiten.Image {
	var err error
	// Return cached image if already loaded
	if images[path] != nil {
		return images[path]
	}
	// Load and cache image
	ext := filepath.Ext(path)
	switch strings.ToLower(ext) {
	case ".jpg":
		images[path], err = LoadJPEG(path)
	case ".jpeg":
		images[path], err = LoadJPEG(path)
	case ".png":
		images[path], err = LoadPNG(path)

	default:
		panic(fmt.Sprintf("failed to load image for extension '%s': %s", ext, path))
	}
	if err != nil {
		panic(err)
	}
	return images[path]
}

func LoadJPEG(path string) (*ebiten.Image, error) {
	background, err := Open(path)
	if err != nil {
		return nil, err
	}
	img, err := JPEGDecode(background)
	if err != nil {
		return nil, err
	}
	return ebiten.NewImageFromImage(img), nil
}

func LoadPNG(path string) (*ebiten.Image, error) {
	file, err := Open(path)
	if err != nil {
		return nil, err
	}
	img, err := PNGDecode(file)
	if err != nil {
		return nil, err
	}
	return ebiten.NewImageFromImage(img), nil
}
