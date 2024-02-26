package leetcode

// dp[i] records the smallest latest value of LIS with length i
func lengthOfLIS(nums []int) int {
	n, len := len(nums), 1
	dp := make([]int, n+1)
	dp[len] = nums[0]

	for _, num := range nums {
		if num > dp[len] {
			len++
			dp[len] = num
		} else {
			var ans, l, r int
			ans = 0
			l, r = 1, len
			for l <= r {
				mid := l + (r-l)>>1
				if dp[mid] < num {
					ans = mid
					l = mid + 1
				} else {
					r = mid - 1
				}
			}
			dp[ans+1] = num
		}
	}

	return len
}
