package leetcode

func trap(height []int) int {
	res := 0
	left := make([]int, len(height))
	right := make([]int, len(height))

	tempMax := 0

	for i, h := range height {
		left[i] = max(0, tempMax-h)
		tempMax = max(tempMax, h)
	}

	tempMax = 0

	for i := len(height) - 1; i >= 0; i-- {
		right[i] = max(0, tempMax-height[i])
		tempMax = max(tempMax, height[i])
	}

	for i := range height {
		res += min(left[i], right[i])
	}

	return res
}
