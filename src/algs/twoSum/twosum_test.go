package main

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	var tests = []struct {
		arr    []int
		target int
		res    []int
		out    []int
	}{
		{[]int{1, 4, 5, 2, 4}, 7, []int{}, []int{2, 3}},
		{[]int{2, 4, -1, 2, 4}, 6, []int{}, []int{0, 4}},
	}

	// index, value
	for _, test := range tests {
		if !reflect.DeepEqual(TwoSum(test.arr, test.target, test.res), test.out) {
			t.Error("Test Failed: expected {} and got {}", test.out, TwoSum(test.arr, test.target, test.res))
		}
	}

}
