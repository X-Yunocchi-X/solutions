package leetcode

func validPartition(nums []int) bool {
	dp := make([]bool, len(nums))

	for i, num := range nums {
		if i == 0 {
			dp[0] = false
			continue
		}
		if num == nums[i-1] && (i-1 == 0 || dp[i-2]) {
			dp[i] = true
			continue
		}
		if i > 1 && num == nums[i-1] && num == nums[i-2] && (i-2 == 0 || dp[i-3]) {
			dp[i] = true
			continue
		}
		if i > 1 && num == nums[i-1]+1 && num == nums[i-2]+2 && (i-2 == 0 || dp[i-3]) {
			dp[i] = true
		}
	}

	return dp[len(dp)-1]
}
