package leetcode

import "slices"

func minimumBoxes(apple []int, capacity []int) int {
	sum := 0
	result := 0
	for _, app := range apple {
		sum += app
	}
	slices.SortFunc(capacity, func(a, b int) int { return b - a })
	for _, capa := range capacity {
		sum -= capa
		result++
		if sum <= 0 {
			break
		}
	}
	return result
}
