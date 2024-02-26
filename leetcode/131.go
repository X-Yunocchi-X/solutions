package leetcode

import "slices"

func partition(s string) [][]string {
	res := [][]string{}
	partion := []string{}
	isValid := func(s string) bool {
		n := len(s)
		for i, j := 0, n-1; i <= j; {
			if s[i] != s[j] {
				return false
			}
			i++
			j--
		}
		return true
	}

	var dfs func(string)
	dfs = func(s string) {
		n := len(s)
		if n == 0 {
			res = append(res, slices.Clone(partion))
			return
		}
		for i := 0; i < n; i++ {
			if isValid(s[:i+1]) {
				partion = append(partion, s[:i+1])
				dfs(s[i+1:])
				partion = partion[:len(partion)-1]
			}
		}
	}
	dfs(s)

	return res
}
