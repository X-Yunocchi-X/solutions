package leetcode

func findKOr(nums []int, k int) int {
	kor := make([]int, 31)
	for _, num := range nums {
		for i := range make([]struct{}, 31) {
			if (1<<i)&num == 1<<i {
				kor[i]++
			}
		}
	}
	result := 0
	for i, bit := range kor {
		if bit >= k {
			result |= (1 << i)
		}
	}
	return result
}
