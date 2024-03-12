package leetcode

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

type FindElements struct {
	set map[int]bool
}

func Constructor(root *TreeNode) FindElements {
	temp := map[int]bool{}
	fe := FindElements{temp}
	dfs(root, true, fe.set)
	return fe
}

func dfs(curr *TreeNode, isRoot bool, set map[int]bool) {
	if isRoot {
		curr.Val = 0
	}
	set[curr.Val] = true
	if curr.Left != nil {
		curr.Left.Val = 2*curr.Val + 1
		dfs(curr.Left, false, set)
	}
	if curr.Right != nil {
		curr.Right.Val = 2*curr.Val + 2
		dfs(curr.Right, false, set)
	}
}

func (this *FindElements) Find(target int) bool {
	return this.set[target]
}
