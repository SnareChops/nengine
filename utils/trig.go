package utils

import (
	"math"

	"github.com/SnareChops/nengine/types"
)

// DistanceBetween returns the distance between two given vectors
func DistanceBetween(start, end types.Position) float64 {
	x1, y1 := start.Pos2()
	x2, y2 := end.Pos2()
	return DistanceBetweenPoints(x1, y1, x2, y2)
}

// DistanceBetweenPoints returns the distance between two given sets of points
func DistanceBetweenPoints(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}

// PointAtAngleWithDistance returns a new (x, y float64) given the starting position
// then traveling at the provided angle the given length distance
func PointAtAngleWithDistance(x, y, angle, dist float64) (float64, float64) {
	return x + dist*math.Cos(angle), y + dist*math.Sin(angle)
}

func AngleBetween(a, b types.Position) float64 {
	x1, y1 := a.Pos2()
	x2, y2 := b.Pos2()
	return AngleBetweenPoints(x1, y1, x2, y2)
}

// AngleBetweenPoints returns the angle from the first point to the second
func AngleBetweenPoints(x1, y1, x2, y2 float64) float64 {
	result := math.Atan2(y2-y1, x2-x1)
	if result < 0 {
		return result + 2*math.Pi
	}
	return result
}

// MoveTowards returns a new (x, y float64) position given the starting position
// then travelling at the provided angle at a consistent speed over time
func MoveTowards(x1, y1, x2, y2, speed float64, delta int) (float64, float64) {
	length := speed / float64(delta)
	dist := DistanceBetweenPoints(x1, y1, x2, y2)
	if dist <= length {
		return x2, y2
	}
	angle := AngleBetweenPoints(x1, y1, x2, y2)
	return PointAtAngleWithDistance(x1, y1, angle, length)
}

// MoveAway returns a new (x, y float64) position given the starting position
// and travelling directly opposite the angle towards the target
func MoveAway(x1, y1, x2, y2, speed float64, delta int) (float64, float64) {
	length := speed / float64(delta)
	angle := AngleBetweenPoints(x1, y1, x2, y2)
	return PointAtAngleWithDistance(x1, y1, angle+math.Pi, length)
}

// Lerp returns the point on a line between the two provided
// endpoints at the given percentage
func Lerp(x1, y1, x2, y2, percent float64) (float64, float64) {
	return x1 + (x2-x1)*percent, y1 + (y2-y1)*percent
}
