package algorithms

// O(n * log(n))
func MergeSort(list []int) []int {
	if len(list) == 1 {
		return list
	}

	middle := len(list) / 2

	a := list[:middle]
	b := list[middle:]

	return merge(MergeSort(a), MergeSort(b))
}

func merge(a, b []int) []int {
	var el int
	merged := []int{}

	i := 0
	j := 0

	for i < len(a) || j < len(b) {
		if i < len(a) && (j >= len(b) || a[i] <= b[j]) {
			el = a[i]
			i++
		} else if j < len(b) {
			el = b[j]
			j++
		}
		merged = append(merged, el)
	}

	return merged
}
