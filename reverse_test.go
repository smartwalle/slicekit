package slicekit_test

import (
	"reflect"
	"testing"

	"github.com/smartwalle/slicekit"
)

func TestReverse(t *testing.T) {
	var tests = []struct {
		name     string
		source   []int
		expected []int
	}{
		{
			name:     "反转多个元素",
			source:   []int{1, 2, 3, 4, 5},
			expected: []int{5, 4, 3, 2, 1},
		},
		{
			name:     "反转偶数个元素",
			source:   []int{1, 2, 3, 4},
			expected: []int{4, 3, 2, 1},
		},
		{
			name:     "单元素切片",
			source:   []int{1},
			expected: []int{1},
		},
		{
			name:     "空切片",
			source:   []int{},
			expected: []int{},
		},
		{
			name:     "nil切片",
			source:   nil,
			expected: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			slicekit.Reverse(test.source)
			if !reflect.DeepEqual(test.source, test.expected) {
				t.Fatalf("实际: %+v, 预期: %+v", test.source, test.expected)
			}
		})
	}
}

func TestReversed(t *testing.T) {
	var source = []int{1, 2, 3, 4}
	var actual = slicekit.Reversed(source)
	var expected = []int{4, 3, 2, 1}
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("实际: %+v, 预期: %+v", actual, expected)
	}
	if !reflect.DeepEqual(source, []int{1, 2, 3, 4}) {
		t.Fatalf("原始切片被修改: %+v", source)
	}

	actual[0] = 100
	if source[3] != 4 {
		t.Fatalf("返回结果与原始切片共享底层数组")
	}
}

func TestReversedEmptySlice(t *testing.T) {
	if actual := slicekit.Reversed([]int{}); actual != nil {
		t.Fatalf("实际: %+v, 预期: nil", actual)
	}
	if actual := slicekit.Reversed([]int(nil)); actual != nil {
		t.Fatalf("实际: %+v, 预期: nil", actual)
	}
}
