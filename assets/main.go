package assets

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/SnareChops/nengine/assets/cache"
	"github.com/SnareChops/nengine/utils"
	"github.com/hajimehoshi/ebiten/v2"
)

func ReloadCache() {
	cache.ReloadCache()
	images = map[string]*ebiten.Image{}
	sheets = map[string]TileSheet{}
}

func AssetsInFolder(path string) (aliases, fonts []string) {
	filepath.Walk("assets/"+path, func(path string, info os.FileInfo, err error) error {
		if info.Mode().IsRegular() && strings.HasSuffix(info.Name(), ".png") {
			aliases = append(aliases, cache.AliasFor(path))
		}
		if info.Mode().IsRegular() && strings.HasSuffix(info.Name(), ".ttf") {
			fonts = append(fonts, cache.FontAliasFor(path))
		}
		return nil
	})
	return
}

func ScaledImage(width, height int, image *ebiten.Image) *ebiten.Image {
	if image.Bounds().Dx() == width && image.Bounds().Dy() == height {
		return image
	}
	result := ebiten.NewImage(width, height)
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Scale(utils.ScaleFactor(image.Bounds().Dx(), image.Bounds().Dy(), width, height))
	result.DrawImage(image, options)
	return result
}
