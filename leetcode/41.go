package leetcode

func firstMissingPositive(nums []int) int {
	res := len(nums)
	for i, num := range nums {
		for num > 0 && num <= len(nums) && nums[num-1] != num {
			nums[i], nums[num-1] = nums[num-1], nums[i]
			num = nums[i]
		}
	}
	for i, num := range nums {
		if i+1 != num {
			res = i + 1
		}
	}
	return res
}
