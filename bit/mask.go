package bit

func IsSet[T ~int](mask T, state T) bool {
	return mask&state == state
}

func BitmaskAdd[T ~int](mask T, flag T) T {
	return mask | flag
}

func BitmaskRemove[T ~int](mask T, flag T) T {
	return mask &^ flag
}
