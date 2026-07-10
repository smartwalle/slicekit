package slicekit

// Reverse 原地反转 slice 中元素的顺序。
func Reverse[T any](slice []T) {
	for left, right := 0, len(slice)-1; left < right; left, right = left+1, right-1 {
		slice[left], slice[right] = slice[right], slice[left]
	}
}

// Reversed 返回一个元素顺序与 slice 相反的新 slice，不会修改原始 slice。
func Reversed[T any](slice []T) []T {
	var n = len(slice)
	if n == 0 {
		return nil
	}
	var reversed = make([]T, n)
	for index, elem := range slice {
		reversed[n-index-1] = elem
	}
	return reversed
}
