package leetcode

import "math/rand"

// counting sort
// func findKthLargest(nums []int, k int) int {
// 	bucket := make([]int, 10000*2+1)
// 	for _, num := range nums {
// 		bucket[num+10000] += 1
// 	}
// 	nowValue := 0
// 	for i := len(bucket) - 1; i >= 0; i-- {
// 		if nowValue+bucket[i] >= k {
// 			return i - 10000
// 		} else {
// 			nowValue += bucket[i]
// 		}
// 	}
// 	return 0
// }

// quickSort
func findKthLargest(nums []int, k int) int {
	return quickSort(nums, 0, len(nums)-1, len(nums)-k)
}

func quickSort(nums []int, start, end, target int) int {
	temp := start + rand.Intn(end-start+1)
	pivot := nums[temp]
	nums[temp], nums[start] = nums[start], nums[temp]
	index := start + 1
	for i := index; i <= end; i++ {
		if nums[i] < pivot {
			nums[i], nums[index] = nums[index], nums[i]
			index++
		}
	}
	nums[start], nums[index-1] = nums[index-1], nums[start]
	index--
	if index == target {
		return pivot
	} else if index > target {
		return quickSort(nums, start, index-1, target)
	} else {
		return quickSort(nums, index+1, end, target)
	}
}
