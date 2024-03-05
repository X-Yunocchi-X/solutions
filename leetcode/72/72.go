package leetcode

func minDistance(word1 string, word2 string) int {
	n, m := len(word1), len(word2)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}
	for i := range dp {
		dp[i][0] = i
	}
	for j := range dp[0] {
		dp[0][j] = j
	}

	for i := range word1 {
		for j := range word2 {
			dp[i+1][j+1] = 10000 // result can not reach
			if word1[i] == word2[j] {
				dp[i+1][j+1] = dp[i][j]
			}
			dp[i+1][j+1] = min(dp[i+1][j+1], min(dp[i][j], dp[i][j+1], dp[i+1][j])+1)
		}
	}

	return dp[n][m]
}
