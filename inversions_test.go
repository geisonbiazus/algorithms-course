package algorithms

import "testing"

func TestSortAndCount(t *testing.T) {
	var testData = []struct {
		list     []int
		expected int
	}{
		{[]int{2}, 0},
		{[]int{2, 1}, 1},
		{[]int{1, 3, 5, 2, 4, 6}, 3},
		{[]int{1, 3, 2, 5, 4, 6}, 2},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 0},
		{[]int{2, 1, 3, 4, 7, 5, 6, 8, 10, 9}, 4},
		{[]int{9, 4, 8, 10, 5, 6, 7, 3, 1, 2}, 34},
	}

	for _, data := range testData {
		_, result := SortAndCount(data.list)
		if result != data.expected {
			t.Errorf("It didn't count the values properly: Expected: %v, Actual: %v", data.expected, result)
		}
	}
}
