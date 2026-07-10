package slicekit_test

import (
	"testing"

	"github.com/smartwalle/slicekit"
)

func TestFind(t *testing.T) {
	var tests = []struct {
		name     string
		source   []int
		expected int
		found    bool
		fn       func(elem int) bool
	}{
		{
			name:     "返回第一个匹配元素",
			source:   []int{1, 2, 4, 6},
			expected: 2,
			found:    true,
			fn: func(elem int) bool {
				return elem%2 == 0
			},
		},
		{
			name:     "匹配零值元素",
			source:   []int{3, 0, 1},
			expected: 0,
			found:    true,
			fn: func(elem int) bool {
				return elem == 0
			},
		},
		{
			name:     "未找到匹配元素",
			source:   []int{1, 2, 3},
			expected: 0,
			found:    false,
			fn: func(elem int) bool {
				return elem > 3
			},
		},
		{
			name:     "空切片",
			source:   nil,
			expected: 0,
			found:    false,
			fn: func(elem int) bool {
				return elem == 1
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var actual, found = slicekit.Find(test.source, test.fn)
			if actual != test.expected || found != test.found {
				t.Fatalf("实际: (%+v, %t), 预期: (%+v, %t)", actual, found, test.expected, test.found)
			}
		})
	}
}

func TestFindStopsAfterFirstMatch(t *testing.T) {
	var calls int
	var actual, found = slicekit.Find([]int{1, 2, 3}, func(elem int) bool {
		calls++
		return elem%2 == 0
	})
	if actual != 2 || !found {
		t.Fatalf("实际: (%+v, %t), 预期: (2, true)", actual, found)
	}
	if calls != 2 {
		t.Fatalf("实际调用次数: %d, 预期调用次数: 2", calls)
	}
}

func TestFindIndex(t *testing.T) {
	var tests = []struct {
		name     string
		source   []int
		expected int
		fn       func(elem int) bool
	}{
		{
			name:     "返回第一个匹配元素索引",
			source:   []int{1, 3, 4, 6},
			expected: 2,
			fn: func(elem int) bool {
				return elem%2 == 0
			},
		},
		{
			name:     "未找到匹配元素",
			source:   []int{1, 2, 3},
			expected: -1,
			fn: func(elem int) bool {
				return elem > 3
			},
		},
		{
			name:     "空切片",
			source:   nil,
			expected: -1,
			fn: func(elem int) bool {
				return elem == 1
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var actual = slicekit.FindIndex(test.source, test.fn)
			if actual != test.expected {
				t.Fatalf("实际: %+v, 预期: %+v", actual, test.expected)
			}
		})
	}
}
