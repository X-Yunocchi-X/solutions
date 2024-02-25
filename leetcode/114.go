/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	curr := root
	for curr != nil {
		if curr.Left == nil {
			curr = curr.Right
			continue
		}
		left := curr.Left
		curr.Left = nil
		right := curr.Right
		curr.Right = left
		rightBottom := curr
		for rightBottom.Right != nil {
			rightBottom = rightBottom.Right
		}
		rightBottom.Right = right
		curr = curr.Right
	}
}
