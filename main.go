package main

import "fmt"

func main() {
	arr := []int{1, 2, -1, 3, -2, 4, 0, 0, 2, -1, 3}
	quickSort(arr)
	fmt.Println(arr)
}

func quickSort(nums []int) {
	n := len(nums)
	if n < 2 {
		return
	}
	pivot := nums[0]
	index := 1
	for i := index; i < n; i++ {
		if nums[i] < pivot {
			nums[i], nums[index] = nums[index], nums[i]
			index++
		}
	}
	nums[index-1], nums[0] = nums[0], nums[index-1]
	quickSort(nums[:index-1])
	quickSort(nums[index:])
}
