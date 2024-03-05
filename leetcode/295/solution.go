package leetcode

import "container/heap"

type minHeap []int

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *minHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *minHeap) Pop() any {
	n := len(*h)
	item := (*h)[n-1]
	*h = (*h)[:n-1]
	return item
}

type maxHeap []int

func (h maxHeap) Len() int           { return len(h) }
func (h maxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h maxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *maxHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *maxHeap) Pop() any {
	n := len(*h)
	item := (*h)[n-1]
	*h = (*h)[:n-1]
	return item
}

type MedianFinder struct {
	minH minHeap
	maxH maxHeap
}

func Constructor() MedianFinder {
	return MedianFinder{minHeap{}, maxHeap{}}
}

func (this *MedianFinder) AddNum(num int) {
	if this.minH.Len() == 0 {
		heap.Push(&this.minH, num)
	} else if this.minH[0] <= num {
		heap.Push(&this.minH, num)
		if this.minH.Len() > this.maxH.Len()+1 {
			heap.Push(&this.maxH, heap.Pop(&this.minH))
		}
	} else {
		heap.Push(&this.maxH, num)
		if this.maxH.Len() > this.minH.Len() {
			heap.Push(&this.minH, heap.Pop(&this.maxH))
		}
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if (this.minH.Len()+this.maxH.Len())%2 == 1 {
		return float64(this.minH[0])
	} else {
		return (float64(this.minH[0]) + float64(this.maxH[0])) / 2
	}
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
