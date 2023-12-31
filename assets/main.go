package assets

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/SnareChops/nengine/assets/cache"
	"github.com/hajimehoshi/ebiten/v2"
)

func ReloadCache() {
	cache.ReloadCache()
	images = map[string]*ebiten.Image{}
	sheets = map[string]TileSheet{}
}

func AssetsInFolder(path string) []string {
	aliases := []string{}
	filepath.Walk("assets/"+path, func(path string, info os.FileInfo, err error) error {
		if info.Mode().IsRegular() && strings.HasSuffix(info.Name(), ".png") {
			aliases = append(aliases, cache.AliasFor(path))
		}
		return nil
	})
	return aliases
}
