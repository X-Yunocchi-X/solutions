package leetcode

func isMatch(s string, p string) bool {
	n, m := len(s), len(p)
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, m)
	}
}
