package assets

import (
	"github.com/SnareChops/nengine/assets/cache"
	"github.com/hajimehoshi/ebiten/v2"
)

type SheetSource struct {
	alias string
	index int
}

func NewSheetSource(alias string, index int) SheetSource {
	return SheetSource{alias, index}
}

func (self SheetSource) Alias() string {
	return self.alias
}

func (self SheetSource) Index() int {
	return self.index
}

func (self SheetSource) Image() *ebiten.Image {
	return ImageFromSheet(self.alias, self.index)
}

type TileSheet struct {
	SheetWidth  int
	SheetHeight int
	CellWidth   int
	CellHeight  int
	Images      []*ebiten.Image
}

var sheets = map[string]TileSheet{}

func GetSheets() map[string]TileSheet {
	return sheets
}

func GetSheet(alias string) TileSheet {
	validateSheetAlias(alias)
	if sheet, ok := sheets[alias]; ok {
		return sheet
	}
	sw, sh, cw, ch, images := cache.PromoteSheet(alias)
	sheets[alias] = TileSheet{
		SheetWidth:  sw,
		SheetHeight: sh,
		CellWidth:   cw,
		CellHeight:  ch,
		Images:      images,
	}
	return sheets[alias]
}

func ImageFromSheet(alias string, idx int) *ebiten.Image {
	sheet := GetSheet(alias)
	return sheet.Images[idx]
}

func validateSheetAlias(alias string) {
	if valid, _, _, _ := cache.IsSheet(alias); !valid {
		panic("Invalid sheet alias: " + alias)
	}
}
