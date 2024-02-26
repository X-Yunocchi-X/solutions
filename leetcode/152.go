package leetcode

func maxProduct(nums []int) int {
	minValue, maxValue, res := nums[0], nums[0], nums[0]

	for i, num := range nums {
		if i == 0 {
			continue
		} else {
			newMinValue := min(num, minValue*num, maxValue*num)
			newMaxValue := max(num, minValue*num, maxValue*num)
			minValue, maxValue = newMinValue, newMaxValue
		}
		res = max(res, maxValue)
	}

	return res
}
