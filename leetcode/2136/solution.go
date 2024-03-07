package leetcode

import "slices"

type Pair struct {
	plant, grow int
}

func earliestFullBloom(plantTime []int, growTime []int) int {
	pairs := make([]Pair, len(plantTime))
	for i := range plantTime {
		pairs[i] = Pair{plantTime[i], growTime[i]}
	}
	slices.SortFunc(pairs, func(a, b Pair) int {
		return a.grow - b.grow
	})
	result := 0
	now := 0
	for _, pair := range pairs {
		temp := now + pair.grow + pair.plant
		result = max(result, temp)
		now += pair.plant
	}
	return result
}
