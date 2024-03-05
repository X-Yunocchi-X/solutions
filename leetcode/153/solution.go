package leetcode

func findMin(nums []int) int {
	result := nums[0]
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) >> 1
		result = min(result, nums[mid])
		if nums[mid] >= nums[l] {
			if nums[r] >= nums[l] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			r = mid - 1
		}
	}
	return result
}
