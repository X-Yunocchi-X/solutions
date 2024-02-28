package leetcode

func jump(nums []int) int {
	maxPosition, end, step := 0, 0, 0
	for i, num := range nums {
		if i == len(nums)-1 {
			break
		}
		maxPosition = max(maxPosition, i+num)
		if i == end {
			step++
			end = maxPosition
		}
	}
	return step
}
