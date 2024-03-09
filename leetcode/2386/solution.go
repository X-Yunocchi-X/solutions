package leetcode

import (
	"container/heap"
	"slices"
)

type Pair struct {
	sum   int64
	index int
}

type minHeap []Pair

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i].sum < h[j].sum }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *minHeap) Push(x any) {
	*h = append(*h, x.(Pair))
}

func (h *minHeap) Pop() any {
	n := len(*h)
	item := (*h)[n-1]
	*h = (*h)[:n-1]
	return item
}

func kSum(nums []int, k int) int64 {
	sum := int64(0)
	for i, num := range nums {
		if num >= 0 {
			sum += int64(num)
		} else {
			nums[i] = -num
		}
	}
	slices.Sort(nums)
	h := minHeap{}
	result := int64(0)
	kth := 0
	heap.Init(&h)
	heap.Push(&h, Pair{0, -1})
	for kth != k {
		kth++
		pop := heap.Pop(&h).(Pair)
		s, i := pop.sum, pop.index
		result = s
		if i == -1 {
			heap.Push(&h, Pair{int64(nums[i+1]), i + 1})
		} else if i+1 < len(nums) {
			heap.Push(&h, Pair{s + int64(nums[i+1]), i + 1})
			heap.Push(&h, Pair{s + int64(nums[i+1]) - int64(nums[i]), i + 1})
		}
	}
	return sum - result
}
