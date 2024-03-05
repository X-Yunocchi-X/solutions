package leetcode

import (
	"container/heap"
	"math"
)

type Pair struct {
	to int
	w  int64
}

type minHeap []Pair

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i].w < h[j].w }
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

func countPaths(n int, roads [][]int) int {
	const mod = 1e9 + 7
	edges := make([][]Pair, n)
	distance := make([]int64, n)
	visited := make([]bool, n)
	ways := make([]int64, n)
	for _, road := range roads {
		f, t, w := road[0], road[1], road[2]
		edges[f] = append(edges[f], Pair{t, int64(w)})
		edges[t] = append(edges[t], Pair{f, int64(w)})
	}
	for i := range make([]struct{}, n) {
		distance[i] = math.MaxInt64
	}
	distance[0] = 0
	ways[0] = 1
	h := minHeap{}
	heap.Init(&h)
	heap.Push(&h, Pair{0, 0})

	for len(h) != 0 {
		pop := heap.Pop(&h).(Pair)
		now, w := pop.to, pop.w
		if visited[now] {
			continue
		}
		for _, edge := range edges[now] {
			to, dis := edge.to, edge.w
			if dis+w < distance[to] {
				ways[to] = ways[now]
				distance[to] = dis + w
				heap.Push(&h, Pair{to, distance[to]})
			} else if dis+w == distance[to] {
				ways[to] += ways[now] % mod
				ways[to] %= mod
			}
		}
	}

	return int(ways[n-1])
}
