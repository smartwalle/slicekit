package slicekit

// Map 将 slice 中的元素转换为 map，fn 返回每个元素对应的键和值。
//
// 当多个元素返回相同的键时，后出现的元素会覆盖先出现的值。
func Map[T any, K comparable, R any](slice []T, fn func(elem T) (K, R)) map[K]R {
	var n = len(slice)
	if n == 0 {
		return nil
	}
	var m = make(map[K]R, n)
	for _, elem := range slice {
		var k, v = fn(elem)
		m[k] = v
	}
	return m
}

// MapMatched 将 slice 中满足 predicate 的元素转换为 map，fn 返回每个元素对应的键和值。
//
// 当多个元素返回相同的键时，后出现的元素会覆盖先出现的值。
func MapMatched[T any, K comparable, R any](slice []T, predicate func(elem T) bool, fn func(elem T) (K, R)) map[K]R {
	var n = len(slice)
	if n == 0 {
		return nil
	}
	var m = make(map[K]R, n)
	for _, elem := range slice {
		if !predicate(elem) {
			continue
		}
		var k, v = fn(elem)
		m[k] = v
	}
	return m
}
