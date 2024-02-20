package utils

type Number interface {
	~int | ~uint | ~float32 | ~float64
}

func Clamp[T Number](num, min, max T) T {
	if num < min {
		return min
	}
	if num > max {
		return max
	}
	return num
}

// ScaleFactor returns the scale factor from one size to another
func ScaleFactor[T Number](fromWidth, fromHeight, toWidth, toHeight T) (float64, float64) {
	return float64(fromWidth) / float64(toWidth), float64(fromHeight) / float64(toHeight)
}

func LinearInterpolate[T ~float32 | ~float64](a, b, percent T) T {
	return a + (b-a)*percent
}
