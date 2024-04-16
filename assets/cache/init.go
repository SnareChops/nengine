package cache

import (
	"image"
	"image/png"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
)

func InitCache() {
	load()
}

func load() {
	images = map[string]image.Image{}
	sheets = map[string]sheet{}
	err := filepath.Walk("assets", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.Mode().IsRegular() {
			if strings.HasSuffix(info.Name(), ".png") {
				file, err := os.Open(path)
				if err != nil {
					return err
				}
				image, err := png.Decode(file)
				if err != nil {
					return err
				}
				if is_sheet, alias, cellWidth, cellHeight := IsSheet(AliasFor(path)); is_sheet {
					addSheet(alias, cellWidth, cellHeight, image)
				} else {
					addImage(AliasFor(path), image)
				}
			}
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
}

func ReloadCache() {
	load()
}

func AliasFor(path string) string {
	path = strings.ReplaceAll(path, "\\", "/")
	regex := regexp.MustCompile(`assets/(.*)\.[png]`)
	matches := regex.FindStringSubmatch(path)
	return matches[1]
}

func FontAliasFor(p string) string {
	p = strings.ReplaceAll(p, "\\", "/")
	return strings.TrimSuffix(path.Base(p), path.Ext(p))
}
