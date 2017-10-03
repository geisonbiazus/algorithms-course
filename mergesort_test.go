package algorithms

import (
	"reflect"
	"testing"
)

var testData = []struct {
	list     []int
	expected []int
}{
	{[]int{2, 1}, []int{1, 2}},
	{[]int{2, 1, 4, 3}, []int{1, 2, 3, 4}},
	{[]int{2, 1, 4}, []int{1, 2, 4}},
	{[]int{1, 2, 4, 5, 6, 7, 8, 9, 10, 11}, []int{1, 2, 4, 5, 6, 7, 8, 9, 10, 11}},
	{[]int{11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0, -1, -2}, []int{-2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}},
}

func TestMergeSort(t *testing.T) {
	for _, data := range testData {
		sorted := MergeSort(data.list)
		if !reflect.DeepEqual(sorted, data.expected) {
			t.Errorf("It didn't sort the values properly: Expected: %v, Actual: %v", data.expected, sorted)
		}
	}
}
