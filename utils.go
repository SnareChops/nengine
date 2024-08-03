package nengine

import (
	"image/color"
	"math"
	"slices"

	"github.com/SnareChops/nengine/types"
	"github.com/SnareChops/nengine/utils"
	"github.com/hajimehoshi/ebiten/v2"
)

// Floats returns a pair of numbers as float64
func Floats[T ~int | ~float32 | ~float64 | ~uint](a, b T) (float64, float64) {
	return float64(a), float64(b)
}

// Ints returns a pair of numbers as int
func Ints[T ~int | ~float32 | ~float64 | ~uint](a, b T) (int, int) {
	return int(a), int(b)
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

func ColorToVec4(color color.Color) [4]float32 {
	r, g, b, a := color.RGBA()
	return [4]float32{
		float32(r) / 0xffff,
		float32(g) / 0xffff,
		float32(b) / 0xffff,
		float32(a) / 0xffff,
	}
}

func GridPointsAroundCell(x, y float64, gridWidth, gridHeight int) []types.Position {
	gw, gh := Floats(gridWidth, gridHeight)
	cx := math.Floor(x)/gw*gw + gw/2
	cy := math.Floor(y)/gh*gh + gh/2
	return []types.Position{
		Point(cx-gw, cy-gh), // Above Left
		Point(cx, cy-gh),    // Above
		Point(cx+gw, cy-gh), // Above Right
		Point(cx+gw, cy),    // Right
		Point(cx+gw, cy+gh), // Below Right
		Point(cx, cy+gh),    // Below
		Point(cx-gw, cy+gh), // Below Left
		Point(cx-gw, cy),    // Left
	}
}

func GridPointsAroundBounds(box Box, gridWidth, gridHeight int) []types.Position {
	gw, gh := Floats(gridWidth, gridHeight)
	x, y := box.Min()

	// snap top left of bounds to grid
	pos := Point(x, y)
	pos.GridAlign(gridWidth, gridHeight)
	// then add half grid width and height to get center
	pos.SetPos2(pos.X()+gw/2, pos.Y()+gh/2)
	// Check if this point is inside the bounds
	if utils.IsPosWithin(box, pos) {
		// If so, subtract full grid cell
		pos.SetPos2(pos.X()-gw, pos.Y()-gh)
	}
	points := []types.Position{pos}
	timesWidth := box.Dx() / gridWidth
	timesHeight := box.Dy() / gridHeight
	// Walk around bounds at grid cell intervals creating points
	for range timesWidth {
		last := points[len(points)-1]
		pos := Point(last.X()+gw, last.Y())
		points = append(points, pos)
	}
	points = append(points, Point(points[len(points)-1].X()+gw, points[len(points)-1].Y()))
	for range timesHeight {
		last := points[len(points)-1]
		pos := Point(last.X(), last.Y()+gh)
		points = append(points, pos)
	}
	points = append(points, Point(points[len(points)-1].X(), points[len(points)-1].Y()+gh))
	for range timesWidth {
		last := points[len(points)-1]
		pos := Point(last.X()-gw, last.Y())
		points = append(points, pos)
	}
	points = append(points, Point(points[len(points)-1].X()-gw, points[len(points)-1].Y()))
	for range timesHeight {
		last := points[len(points)-1]
		pos := Point(last.X(), last.Y()-gh)
		points = append(points, pos)
	}
	return points
}

func ClosestPointAroundBounds(bounds Bounds, origin types.Position, gridWidth, gridHeight int) types.Position {
	points := GridPointsAroundBounds(bounds, gridWidth, gridHeight)
	if len(points) == 0 {
		return nil
	}
	slices.SortFunc(points, func(a, b types.Position) int {
		return int(DistanceBetween(origin, a) - DistanceBetween(origin, b))
	})
	return points[0]
}
