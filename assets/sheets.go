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
	Alias       string
	SheetWidth  int
	SheetHeight int
	CellWidth   int
	CellHeight  int
	Images      []*ebiten.Image
}

// RowLen returns the number of cells in a row (left-to-right)
func (self TileSheet) RowLen() int {
	return self.SheetWidth / self.CellWidth
}

// ColLen returns the number of cells in a column (top-to-bottom)
func (self TileSheet) ColLen() int {
	return self.SheetHeight / self.CellHeight
}

func (self TileSheet) Sources() []SheetSource {
	sources := []SheetSource{}
	for i := range self.Images {
		sources = append(sources, NewSheetSource(self.Alias, i))
	}
	return sources
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
		Alias:       alias,
		SheetWidth:  sw,
		SheetHeight: sh,
		CellWidth:   cw,
		CellHeight:  ch,
		Images:      images,
	}
	return sheets[alias]
}

// GetSheetRange returns a slice of images within the range [start:end] from the sheet
func GetSheetRange(alias string, start, end int) []*ebiten.Image {
	sheet := GetSheet(alias)
	return sheet.Images[start:end]
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
