package leetcode

import "math"

func removeDuplicates(nums []int) int {
	now := math.MinInt
	invalid := 0
	for i, num := range nums {
		if i < 2 {
			continue
		}
		if num == nums[i-1] && num == nums[i-2] {
			now = num
		}
		if now == num {
			nums[i] = math.MinInt
			invalid++
		}
	}
	i, j := 0, 0
	for j < len(nums) {
		for j < len(nums) && nums[j] != math.MinInt {
			j++
		}
		if j == len(nums) {
			break
		}
		i = j
		for i < len(nums) && nums[i] == math.MinInt {
			i++
		}
		if i == len(nums) {
			break
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	return len(nums) - invalid
}
