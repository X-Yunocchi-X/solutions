package leetcode

func canPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum&1 == 1 {
		return false
	}

	target := sum >> 1
	dp := make([]bool, target+1)
	dp[0] = true

	for _, num := range nums {
		for j := target; j >= 0; j-- {
			if dp[j] && j+num <= target {
				dp[j+num] = true
			}
		}
	}

	return dp[target]
}
