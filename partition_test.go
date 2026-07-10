package slicekit_test

import (
	"testing"

	"github.com/smartwalle/slicekit"
)

func TestPartition(t *testing.T) {
	var tests = []struct {
		name      string
		source    []int
		matched   []int
		unmatched []int
		fn        func(elem int) bool
	}{
		{
			name:      "按奇偶数分组",
			source:    []int{1, 2, 3, 4, 5, 6},
			matched:   []int{2, 4, 6},
			unmatched: []int{1, 3, 5},
			fn: func(elem int) bool {
				return elem%2 == 0
			},
		},
		{
			name:      "所有元素均匹配",
			source:    []int{1, 2, 3},
			matched:   []int{1, 2, 3},
			unmatched: []int{},
			fn: func(elem int) bool {
				return true
			},
		},
		{
			name:      "没有元素匹配",
			source:    []int{1, 2, 3},
			matched:   []int{},
			unmatched: []int{1, 2, 3},
			fn: func(elem int) bool {
				return false
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var matched, unmatched = slicekit.Partition(test.source, test.fn)
			if !slicekit.Equals(matched, test.matched, IntEqual) {
				t.Fatalf("匹配结果实际: %+v, 预期: %+v", matched, test.matched)
			}
			if !slicekit.Equals(unmatched, test.unmatched, IntEqual) {
				t.Fatalf("未匹配结果实际: %+v, 预期: %+v", unmatched, test.unmatched)
			}
		})
	}
}

func TestPartitionEmptySlice(t *testing.T) {
	var matched, unmatched = slicekit.Partition([]int(nil), func(elem int) bool {
		return elem > 0
	})
	if matched != nil || unmatched != nil {
		t.Fatalf("实际: (%+v, %+v), 预期: (nil, nil)", matched, unmatched)
	}
}
