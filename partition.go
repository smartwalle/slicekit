package slicekit

// Partition 将 slice 中的元素按指定条件分为两组。
//
// 第一个返回值包含满足条件的元素，第二个返回值包含不满足条件的元素；两组元素均保持原始顺序。
func Partition[T any](slice []T, fn func(elem T) bool) ([]T, []T) {
	var n = len(slice)
	if n == 0 {
		return nil, nil
	}
	var matched = make([]T, 0, n)
	var unmatched = make([]T, 0, n)
	for _, elem := range slice {
		if fn(elem) {
			matched = append(matched, elem)
			continue
		}
		unmatched = append(unmatched, elem)
	}
	return matched, unmatched
}
