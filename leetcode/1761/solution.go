package leetcode

func minTrioDegree(n int, edges [][]int) int {
	result := -1
	paths := make([][]int, n)
	degree := make([]int, n)
	for i := range paths {
		paths[i] = make([]int, n)
	}
	for _, edge := range edges {
		from, to := edge[0]-1, edge[1]-1
		paths[from][to] = 1
		paths[to][from] = 1
		degree[from]++
		degree[to]++
	}
	for i := range paths {
		for j := i + 1; j < n; j++ {
			if paths[i][j] != 1 {
				continue
			}
			for k := j + 1; k < n; k++ {
				if paths[i][k] == 1 && paths[j][k] == 1 {
					temp := degree[i] + degree[j] + degree[k] - 6
					if result == -1 {
						result = temp
					} else {
						result = min(result, temp)
					}
				}
			}
		}
	}
	return result
}
