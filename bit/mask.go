package bit

func IsSet[T ~int | ~byte](mask T, state T) bool {
	return mask&state == state
}

func BitmaskAdd[T ~int | ~byte](mask T, flag T) T {
	return mask | flag
}

func BitmaskRemove[T ~int | ~byte](mask T, flag T) T {
	return mask &^ flag
}
