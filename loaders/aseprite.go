package loaders

import (
	"errors"
	"fmt"

	"github.com/SnareChops/aseprite-loader/lib"
	"github.com/hajimehoshi/ebiten/v2"
)

func PreloadImageAseprite(alias, path string) {
	if _, ok := flat[alias]; ok {
		return
	}
	frames, err := lib.LoadFrames(path)
	if err != nil {
		panic(fmt.Errorf("PreloadImageAseprite: %s\n%s", path, err))
	}
	if len(frames) != 1 {
		panic(errors.New("tried to load a flat image from a multi-frame file." + path))
	}
	flat[alias] = ebiten.NewImageFromImage(frames[0].Image)
}

func PreloadSheetAseprite(alias, path string) {
	if _, ok := sheets[alias]; ok {
		return
	}
	frames, err := lib.LoadFrames(path)
	if err != nil {
		panic(fmt.Errorf("PreloadSheetAseprite: %s\n%s", path, err))
	}
	if len(frames) != 1 {
		panic(errors.New("tried to load a sheet from a multi-frame file." + path))
	}
	slices, err := lib.Slice(frames[0].Image, frames[0].GridWidth, frames[0].GridHeight)
	if err != nil {
		panic(fmt.Errorf("PreloadSheetAseprite: %s\n%s", path, err))
	}
	cells := []*ebiten.Image{}
	for _, slice := range slices {
		cells = append(cells, ebiten.NewImageFromImage(slice))
	}
	sheets[alias] = Sheet{
		CellWidth:  frames[0].GridWidth,
		CellHeight: frames[0].GridHeight,
		Cells:      cells,
	}
}

func PreloadAnimAseprite(alias, path string) {
	if _, ok := anims[alias]; ok {
		return
	}
	frames, err := lib.LoadFrames(path)
	if err != nil {
		panic(fmt.Errorf("PreloadAnimAseprite: %s\n%s", path, err))
	}
	cells := []*ebiten.Image{}
	for _, frame := range frames {
		cells = append(cells, ebiten.NewImageFromImage(frame.Image))
	}
	anims[alias] = Anim{
		Duration:    frames[0].Duration,
		FrameWidth:  frames[0].GridWidth,
		FrameHeight: frames[0].GridHeight,
		Frames:      cells,
	}
}
