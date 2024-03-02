package main

import (
	"fmt"

	"github.com/X-Yunocchi-X/Solution/leetcode"
)

func main() {
	res := leetcode.TopKFrequent([]int{6, 0, 1, 4, 9, 7, -3, 1, -4, -8, 4, -7, -3, 3, 2, -3, 9, 5, -4, 0}, 6)
	fmt.Println(res)
}
