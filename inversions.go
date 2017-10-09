package algorithms

// O(n * log(n))
// SortAndCount counts the number of inversions by sorting an array
// Example: 1, 3, 5, 2, 4, 6
//                2, 5
//             2, 3
//                   4, 5
//          1, 2, 3, 4, 5, 6
// This exaple performs 3 invertions to sort the array
func SortAndCount(list []int) ([]int, int) {
	if len(list) <= 1 {
		return list, 0
	}

	middle := len(list) / 2

	a, invA := SortAndCount(list[:middle])
	b, invB := SortAndCount(list[middle:])
	c, invC := sortAndCountSplitInv(a, b)

	return c, invA + invB + invC
}

func sortAndCountSplitInv(a, b []int) ([]int, int) {
	el := 0
	inv := 0
	merged := []int{}

	i := 0
	j := 0

	for i < len(a) || j < len(b) {
		if i < len(a) && (j >= len(b) || a[i] <= b[j]) {
			el = a[i]
			i++
		} else if j < len(b) {
			inv = inv + len(a) - i
			el = b[j]
			j++
		}
		merged = append(merged, el)
	}

	return merged, inv
}
