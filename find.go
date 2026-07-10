package slicekit

// Find 返回 slice 中第一个满足指定条件的元素。
//
// 当没有元素满足条件时，返回元素类型的零值和 false。
func Find[T any](slice []T, fn func(elem T) bool) (T, bool) {
	for _, elem := range slice {
		if fn(elem) {
			return elem, true
		}
	}
	var elem T
	return elem, false
}

// FindIndex 返回 slice 中第一个满足指定条件的元素索引。
//
// 当没有元素满足条件时，返回 -1。
func FindIndex[T any](slice []T, fn func(elem T) bool) int {
	for index, elem := range slice {
		if fn(elem) {
			return index
		}
	}
	return -1
}
