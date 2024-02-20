package nengine

import "github.com/hajimehoshi/ebiten/v2"

// Floats returns a pair of int as float64
func Floats(a, b int) (float64, float64) {
	return float64(a), float64(b)
}

// Ints returns a pair of float64 as int
func Ints(a, b float64) (int, int) {
	return int(a), int(b)
}

func IsSet[T ~int](mask T, state T) bool {
	return mask&state == state
}

func RelativePosition[T ~int | ~float64](x, y T, bounds Bounds) (T, T) {
	bx, by := bounds.Min()
	return x - T(bx), y - T(by)
}

type verticalRelative interface {
	Pos2() (float64, float64)
	Dy() int
}

func PositionBelow(object verticalRelative, padding int) (float64, float64) {
	x, y := object.Pos2()
	return x, y + float64(object.Dy()+padding)
}

type horizontalRelative interface {
	Pos2() (float64, float64)
	Dx() int
}

func PositionRight(object horizontalRelative, padding int) (float64, float64) {
	x, y := object.Pos2()
	return x + float64(object.Dx()+padding), y
}

func FitToNewImage(w, h int, image *ebiten.Image) *ebiten.Image {
	out := ebiten.NewImage(w, h)
	wf, hf := ScaleFactor(image.Bounds().Dx(), image.Bounds().Dy(), w, h)
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Scale(wf, hf)
	out.DrawImage(image, options)
	return out
}
