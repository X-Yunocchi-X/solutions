/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
package leetcode

func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	result := 0
	sum := map[int64]int{0: 1}
	var dfs func(*TreeNode, int)
	dfs = func(curr *TreeNode, val int) {
		if curr == nil {
			return
		}
		val += curr.Val
		result += sum[int64(val)-int64(targetSum)]
		sum[int64(val)]++
		dfs(curr.Left, val)
		dfs(curr.Right, val)
		sum[int64(val)]--
	}
	dfs(root, 0)
	return result
}
