package main

import "fmt"

func main() {
	var nums []int
	if nums == nil {
		fmt.Println("true")
	}
	result := append(nums, 1, 2, 3)
	fmt.Println(result)
}
