package leetcode

func generateParenthesis(n int) []string {
	result := []string{}
	left, right := 0, 0
	tempStr := []rune{}
	var dfs func()

	dfs = func() {
		if left == n && right == n {
			result = append(result, string(tempStr))
			return
		}
		if left < n {
			left++
			tempStr = append(tempStr, '(')
			dfs()
			tempStr = tempStr[:len(tempStr)-1]
			left--
		}
		if left > right {
			right++
			tempStr = append(tempStr, ')')
			dfs()
			tempStr = tempStr[:len(tempStr)-1]
			right--
		}
	}
	dfs()

	return result
}
