package leetcode

import "container/heap"

type minHeap [][2]int // num, count

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i][1] < h[j][1] }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *minHeap) Push(x any) {
	*h = append(*h, x.([2]int))
}

func (h *minHeap) Pop() any {
	n := len(*h)
	item := (*h)[n-1]
	*h = (*h)[:n-1]
	return item
}

func topKFrequent(nums []int, k int) []int {
	h := minHeap{}
	heap.Init(&h)
	result := []int{}
	counts := map[int]int{}
	for _, num := range nums {
		counts[num]++
	}
	for num, count := range counts {
		if len(h) < k {
			heap.Push(&h, [2]int{num, count})
		} else if h[0][1] < count {
			heap.Pop(&h)
			heap.Push(&h, [2]int{num, count})
		}
	}
	for len(h) != 0 {
		result = append(result, heap.Pop(&h).([2]int)[0])
	}
	return result
}
