package bit

func IsSet[T ~int | ~uint | ~byte | ~uint64](mask T, state T) bool {
	return mask&state == state
}

func BitmaskAdd[T ~int | ~uint | ~byte | ~uint64](mask T, flag T) T {
	return mask | flag
}

func BitmaskRemove[T ~int | ~uint | ~byte | ~uint64](mask T, flag T) T {
	return mask &^ flag
}
