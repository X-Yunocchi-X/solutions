package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
	result := root.Val
	var dfs func(*TreeNode) int
	dfs = func(curr *TreeNode) int {
		if curr == nil {
			return 0
		}
		left, right := dfs(curr.Left), dfs(curr.Right)
		value := curr.Val + max(0, left, right)
		result = max(result, value, left+right+curr.Val)
		return value
	}
	dfs(root)
	return result
}
