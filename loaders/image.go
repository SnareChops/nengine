package loaders

import (
	"fmt"
	"strings"

	"github.com/SnareChops/aseprite-loader/lib"
	"github.com/hajimehoshi/ebiten/v2"
)

type SheetSource struct {
	alias string
	index int
}

func (s SheetSource) Alias() string {
	return s.alias
}
func (s SheetSource) Index() int {
	return s.index
}
func (s SheetSource) Image() *ebiten.Image {
	return GetSheetCell(s.alias, s.index)
}

type Sheet struct {
	Alias      string
	Width      int
	Height     int
	CellWidth  int
	CellHeight int
	Cells      []*ebiten.Image
}

func (self Sheet) Sources() []SheetSource {
	sources := []SheetSource{}
	for i := range self.Cells {
		sources = append(sources, SheetSource{self.Alias, i})
	}
	return sources
}

type Anim struct {
	Duration    int
	Tags        []lib.Tag
	FrameWidth  int
	FrameHeight int
	Frames      []*ebiten.Image
}

var flat = map[string]*ebiten.Image{}
var sheets = map[string]Sheet{}
var anims = map[string]Anim{}

func imageType(path string) string {
	p := strings.ToLower(path)
	if strings.HasSuffix(p, ".aseprite") || strings.HasSuffix(p, ".ase") {
		return "aseprite"
	}
	if strings.HasSuffix(p, ".png") {
		return "png"
	}
	if strings.HasSuffix(p, ".jpg") || strings.HasSuffix(p, ".jpeg") {
		return "jpeg"
	}
	panic("imageType: Unknown image type: " + path)
}

func PreloadImage(alias, path string) {
	switch imageType(path) {
	case "aseprite":
		PreloadImageAseprite(alias, path)
	case "png":
		PreloadImagePng(alias, path)
	case "jpeg":
		PreloadImageJpeg(alias, path)
	default:
		panic("PreloadImage: Unsupported image type" + alias)
	}
}

func PreloadSheet(alias, path string) {
	switch imageType(path) {
	case "aseprite":
		PreloadSheetAseprite(alias, path)
	case "png":
		PreloadSheetPng(alias, path)
	case "jpeg":
		PreloadSheetJpeg(alias, path)
	default:
		panic("PreloadSheet: Unsupported image type" + alias)
	}
}

func PreloadAnim(alias, path string) {
	switch imageType(path) {
	case "aseprite":
		PreloadAnimAseprite(alias, path)
	case "png":
		panic("PreloadAnim: PNG not supported for animations. Use PreloadSheet instead.")
	case "jpeg":
		panic("PreloadAnim: JPG not supported for animations. Use PreloadSheet instead.")
	default:
		panic("PreloadAnim: Unsupported image type" + alias)
	}
}

func GetImage(alias string) *ebiten.Image {
	if image, ok := flat[alias]; ok {
		return image
	}
	panic("GetImage: " + alias + " not found in cache. Did you forget to Preload?")
}

func GetSheet(alias string) Sheet {
	if sheet, ok := sheets[alias]; ok {
		return sheet
	}
	panic("GetSheet: " + alias + " not found in cache. Did you forget to Preload?")
}

func GetSheetCell(alias string, index int) *ebiten.Image {
	sheet := GetSheet(alias)
	if index < 0 || index >= len(sheet.Cells) {
		panic(fmt.Sprintf("GetSheetCell:%d: %s out of range", index, alias))
	}
	return sheet.Cells[index]
}

func GetSheetRange(alias string, start, end int) []*ebiten.Image {
	sheet := GetSheet(alias)
	if start < 0 || end >= len(sheet.Cells) {
		panic(fmt.Sprintf("GetSheetRange:%d-%d: %s out of range", start, end, alias))
	}
	return sheet.Cells[start:end]
}

func GetAnim(alias string) Anim {
	if anim, ok := anims[alias]; ok {
		return anim
	}
	panic("GetAnim: " + alias + " not found in cache. Did you forget to Preload?")
}
