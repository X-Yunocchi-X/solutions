package leetcode

func maxArrayValue(nums []int) int64 {
	n := len(nums)
	result := int64(0)
	if n == 1 {
		return int64(nums[0])
	}
	for i := n - 2; i >= 0; i-- {
		if nums[i] <= nums[i+1] {
			nums[i] += nums[i+1]
		}
		result = max(result, int64(nums[i]))
	}
	return result
}
