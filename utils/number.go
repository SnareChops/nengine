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

func LinearInterpolate[T ~float32 | ~float64](a, b, percent T) T {
	return a + (b-a)*percent
}
