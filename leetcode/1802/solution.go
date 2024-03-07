package leetcode

func maxValue(n int, index int, maxSum int) int {
	ans := maxSum
	l, r := 1, maxSum
	calculate := func(m, length int) int {
		result := 0
		if m > length {
			result = length * (2*m - 1 - length) / 2
		} else {
			result = m*(m-1)/2 + length - m + 1
		}
		return result
	}
	check := func(m int) bool {
		result := m + calculate(m, index) + calculate(m, n-index-1)
		return result <= maxSum
	}
	for l <= r {
		mid := l + (r-l)>>1
		if check(mid) {
			ans = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return ans
}
