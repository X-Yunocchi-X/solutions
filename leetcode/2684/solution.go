package leetcode

func maxMoves(grid [][]int) int {
	result := 0
	n, m := len(grid), len(grid[0])
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, m)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	var dfs func(int, int) int
	movement := [3][2]int{{-1, 1}, {0, 1}, {1, 1}}
	dfs = func(row, col int) int {
		if dp[row][col] != -1 {
			return dp[row][col]
		}
		result := 1
		for _, move := range movement {
			tox, toy := row+move[0], col+move[1]
			if tox != -1 && toy != m && tox != n && grid[tox][toy] > grid[row][col] {
				result = max(result, dfs(tox, toy)+1)
			}
		}
		dp[row][col] = result
		return result
	}
	for i := range grid {
		result = max(result, dfs(i, 0))
	}
	return result - 1
}
