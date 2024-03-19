package leetcode

func findMinHeightTrees(n int, edges [][]int) []int {
	tree := make([][]int, n)
	for _, edge := range edges {
		from, to := edge[0], edge[1]
		tree[from] = append(tree[from], to)
		tree[to] = append(tree[to], from)
	}
	var dfs func(int, int, int) int
	result := []int{}
	minDepth := -1
	dfs = func(curr, fa, now int) int {
		depth := 1
		for _, to := range tree[curr] {
			if to == fa {
				continue
			}
			depth = max(depth, dfs(to, curr, now+1)+1)
		}
		tempDepth := max(depth, now)
		if minDepth == -1 || tempDepth < minDepth {
			minDepth = tempDepth
			result = []int{curr}
		} else if tempDepth == minDepth {
			result = append(result, curr)
		}
		return depth
	}
	dfs(0, -1, 0)
	return result
}
