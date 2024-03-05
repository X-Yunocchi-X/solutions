package leetcode

type minHeap [][2]int // index, value

// func TopKFrequent(nums []int, k int) []int {
// 	result := []int{}
// 	heap := make(minHeap, 0)
// 	// heap.push([2]int{3333, 1})
// 	// heap.push([2]int{1, 3})
// 	mm := map[int]int{}
// 	for _, num := range nums {
// 		mm[num]++
// 	}
// 	for index, num := range mm {
// 		if heap.len() < k {
// 			heap.push([2]int{index, num})
// 		} else if heap.peek()[1] < num {
// 			heap.pop()
// 			heap.push([2]int{index, num})
// 		}
// 	}
// 	for heap.len() != 0 {
// 		result = append(result, heap.pop()[0])
// 	}
// 	return result
// }
