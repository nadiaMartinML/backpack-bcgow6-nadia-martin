package sort

import "sort"

func SortingAsc(num []int) []int {
	sort.Slice(num, func(i, j int) bool {
		return num[i] < num[j]
	})
	return num
}
