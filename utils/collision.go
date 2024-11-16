package utils

import "github.com/SnareChops/nengine/types"

func DoesCollide(a, b types.Box) bool {
	x1m, y1m := a.Min()
	x1M, y1M := a.Max()

	x2m, y2m := b.Min()
	x2M, y2M := b.Max()

	return !(x2M < x1m || x2m > x1M || y2M < y1m || y2m > y1M)
}

func IsWithin[T ~int | ~float64](box types.Box, x, y T) bool {
	w, h := box.Size()
	bx, by := box.Min()
	if w == 1 && h == 1 {
		return float64(x) == bx && float64(y) == by
	}
	x2, y2 := box.Max()
	return float64(x) >= bx && float64(x) <= x2 && float64(y) >= by && float64(y) <= y2
}

func IsPosWithin(box types.Box, pos types.Position) bool {
	return IsWithin(box, pos.X(), pos.Y())
}

func RawCollide[T ~int | ~float64](x1, y1, w1, h1, x2, y2, w2, h2 T) bool {
	x1M := x1 + w1
	y1M := y1 + h1

	x2M := x2 + w2
	y2M := y2 + h2

	return !(x2M < x1 || x2 > x1M || y2M < y1 || y2 > y1M)
}
