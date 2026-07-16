package slicekit

// Apply 对 slice 中的每个元素应用 fn，并返回转换后的 slice。
func Apply[T any, R any](slice []T, fn func(elem T) R) []R {
	var n = len(slice)
	if n == 0 {
		return nil
	}
	var ns = make([]R, n)
	for idx, elem := range slice {
		ns[idx] = fn(elem)
	}
	return ns
}

// ApplyMatched 对 slice 中满足指定条件的元素应用 fn，并返回转换后的 slice。
func ApplyMatched[T any, R any](slice []T, predicate func(elem T) bool, fn func(elem T) R) []R {
	var n = len(slice)
	if n == 0 {
		return nil
	}
	var ns = make([]R, 0, n)
	for _, elem := range slice {
		if predicate(elem) {
			ns = append(ns, fn(elem))
		}
	}
	return ns
}
