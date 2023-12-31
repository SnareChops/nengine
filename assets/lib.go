package assets

import (
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"path/filepath"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	// OS
	ReadFile = os.ReadFile
	Open     = func(path string) (io.Reader, error) {
		root, err := os.Getwd()
		if err != nil {
			return nil, err
		}
		return os.Open(filepath.Join(root, path))
	}

	// Images
	JPEGDecode        = jpeg.Decode
	PNGDecode         = png.Decode
	NewImage          = ebiten.NewImage
	NewImageFromImage = ebiten.NewImageFromImage
)
