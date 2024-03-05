package leetcode

func longestValidParentheses(s string) int {
	dp := make([][]int, len(s)) // length, index
	result := 0
	for i, char := range s {
		if i == 0 || char == '(' {
			continue
		}
		if s[i-1] == '(' {
			if i-2 >= 0 && dp[i-2] != nil {
				dp[i] = []int{dp[i-2][0] + 2, dp[i-2][1]}
			} else {
				dp[i] = []int{2, i - 1}
			}
		} else {
			if dp[i-1] != nil && dp[i-1][1] != 0 {
				if s[dp[i-1][1]-1] == '(' {
					if dp[i-1][1]-2 >= 0 && dp[dp[i-1][1]-2] != nil {
						dp[i] = []int{dp[i-1][0] + 2 + dp[dp[i-1][1]-2][0], dp[dp[i-1][1]-2][1]}
					} else {
						dp[i] = []int{dp[i-1][0] + 2, dp[i-1][1] - 1}
					}
				}
			}
		}
		if dp[i] != nil {
			result = max(result, dp[i][0])
		}
	}
	return result
}
