package leetcode

import "slices"

func maximumHappinessSum(happiness []int, k int) int64 {
	result := int64(0)
	slices.SortFunc(happiness, func(a, b int) int { return b - a })
	for i, happy := range happiness {
		if i == k {
			break
		}
		result += int64(max(0, happy-i))
	}
	return result
}
