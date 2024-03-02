package leetcode

func largestRectangleArea(heights []int) int {
	result := 0
	left, right := make([]int, len(heights)), make([]int, len(heights))
	for i := range heights {
		left[i] = -1
		right[i] = -1
	}
	stack := []int{}
	for i, height := range heights {
		for len(stack) != 0 && heights[stack[len(stack)-1]] > height {
			pop := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			right[pop] = i
		}
		if len(stack) != 0 {
			left[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}
	for i := range heights {
		l, r := left[i], right[i]
		if l == -1 {
			l = i
		} else {
			l = i - l
		}
		if r == -1 {
			r = len(heights) - i - 1
		} else {
			r = r - i
		}
		result = max(result, heights[i]*(l+r+1))
	}
	return result
}
