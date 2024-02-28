package leetcode

func longestCommonSubsequence(text1 string, text2 string) int {
	len1, len2 := len(text1), len(text2)
	dp := make([][]int, len1+1)
	for i := range dp {
		dp[i] = make([]int, len2+1)
	}
	for i := range text1 {
		for j := range text2 {
			if text1[i] == text2[j] {
				dp[i+1][j+1] = dp[i][j] + 1
			}
			dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j], dp[i+1][j+1])
		}
	}
	return dp[len1][len2]
}
