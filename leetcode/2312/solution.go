package leetcode

func sellingWood(m int, n int, prices [][]int) int64 {
	dp := make([][]int, m+1)
	priceMap := make([][]int, m+1)
	for i := range priceMap {
		priceMap[i] = make([]int, n+1)
	}
	for _, price := range prices {
		priceMap[price[0]][price[1]] = price[2]
	}
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	var dfs func(int, int) int
	for i := range dp {
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	dfs = func(i, j int) int {
		if i < 0 || j < 0 || i > m || j > n {
			return 0
		}
		if dp[i][j] != -1 {
			return dp[i][j]
		}
		dp[i][j] = priceMap[i][j]
		for k := 1; k < i; k++ {
			dp[i][j] = max(dp[i][j], dfs(i-k, j)+dfs(k, j))
		}
		for k := 1; k < j; k++ {
			dp[i][j] = max(dp[i][j], dfs(i, j-k)+dfs(i, k))
		}
		return dp[i][j]
	}
	dfs(m, n)
	return int64(dp[m][n])
}
