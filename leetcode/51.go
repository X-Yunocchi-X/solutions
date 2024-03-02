package leetcode

func solveNQueens(n int) [][]string {
	result := [][]string{}
	solveMap := make([][]int, n)
	for i := range solveMap {
		solveMap[i] = make([]int, n)
	}
	lines := make([]int, n)
	diagonal1, diagonal2 := make([]int, 2*n-1), make([]int, 2*n-1)
	var dfs func(int)
	dfs = func(now int) {
		if now == n {
			temp := []string{}
			for _, line := range solveMap {
				str := ""
				for _, num := range line {
					if num == 0 {
						str += "."
					} else {
						str += "Q"
					}
				}
				temp = append(temp, str)
			}
			result = append(result, temp)
			return
		}
		for i := range make([]int, n) {
			if lines[i] == 1 || diagonal1[now+i] == 1 || diagonal2[now+(n-i)-1] == 1 {
				continue
			}
			solveMap[now][i] = 1
			lines[i] = 1
			diagonal1[now+i] = 1
			diagonal2[now+(n-i)-1] = 1
			dfs(now + 1)
			solveMap[now][i] = 0
			lines[i] = 0
			diagonal1[now+i] = 0
			diagonal2[now+(n-i)-1] = 0
		}
	}
	dfs(0)
	return result
}
