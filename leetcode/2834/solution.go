package leetcode

func minimumPossibleSum(n int, target int) int {
	result := 0
	mod := int(1e9 + 7)
	m := target / 2
	if n <= m {
		result = n * (n + 1) / 2 % mod
	} else {
		result = m*(m+1)/2%mod + (n-m)*(target+target+n-m-1)/2%mod
	}
	return result % mod
}
