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
	list := map[*TreeNode]map[int]int{}
	result := 0

	var dfs func(*TreeNode, int)
	dfs = func(root *TreeNode, targetSum int) {
		if root == nil {
			return
		}
		list[root] = map[int]int{}
		dfs(root.Left, targetSum)
		dfs(root.Right, targetSum)
		if num, ok := list[root][root.Val]; ok {
			list[root][root.Val] = num + 1
		}
		if root.Left != nil {
			for key, value := range list[root.Left] {
				if num, ok := list[root][key]; ok {
					list[root][key+root.Val] = num + value
				} else {
					list[root][key+root.Val] = value
				}
			}
		}
		if root.Right != nil {
			for key, value := range list[root.Right] {
				if num, ok := list[root][key]; ok {
					list[root][key+root.Val] = num + value
				} else {
					list[root][key+root.Val] = value
				}
			}
		}
		if num, ok := list[root][targetSum]; ok {
			result += num
		}
	}
	dfs(root, targetSum)
	return result
}
