package leetcode

func reachableNodes(n int, edges [][]int, restricted []int) int {
	var dfs func(int) int
	edgeSet := make([][]int, n)
	restrictSet := make([]bool, n)
	for _, res := range restricted {
		restrictSet[res] = true
	}
	for _, edge := range edges {
		from, to := edge[0], edge[1]
		edgeSet[from] = append(edgeSet[from], to)
		edgeSet[to] = append(edgeSet[to], from)
	}
	visited := make([]bool, n)
	dfs = func(curr int) int {
		if visited[curr] || restrictSet[curr] {
			return 0
		}
		result := 1
		visited[curr] = true
		for _, to := range edgeSet[curr] {
			result += dfs(to)
		}
		return result
	}
	return dfs(0)
}
