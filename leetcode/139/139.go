package leetcode

func wordBreak(s string, wordDict []string) bool {
	n := len(s)
	dp := make([]bool, n)

	for i := range s {
		for _, word := range wordDict {
			if i+1 < len(word) {
				continue
			}
			k := len(word)
			if s[i-k+1:i+1] == word {
				if i-k == -1 || dp[i-k] {
					dp[i] = true
				}
			}
		}
	}

	return dp[n-1]
}
