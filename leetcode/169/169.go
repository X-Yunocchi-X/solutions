package leetcode

func majorityElement(nums []int) int {
	curr, count := 0, 0
	for _, num := range nums {
		if count == 0 {
			curr = num
			count = 1
			continue
		}
		if num == curr {
			count++
		} else {
			count--
		}
	}
	return curr
}
