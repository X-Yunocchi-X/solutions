package leetcode

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	result := -1
	for l <= r {
		mid := l + (r-l)>>1
		if nums[mid] == target {
			result = mid
			break
		}
		if nums[mid] >= nums[l] {
			if nums[mid] >= target && nums[l] <= target {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			if nums[mid] <= target && nums[r] >= target {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return result
}
