package leetcode

import "slices"

func partitionLabels(s string) []int {
	set := map[int]struct{}{}
	result := []int{}
	firstOccur := make([]int, 26)
	for i, char := range s {
		if _, ok := set[int(char)]; !ok {
			set[int(char)] = struct{}{}
			firstOccur[int(char)-'a'] = i
		}
	}
	clear(set)

	position := len(s) - 1
	end := position
	for i := len(s) - 1; i >= 0; i-- {
		position = min(position, firstOccur[int(s[i])-'a'])
		if i == position {
			result = append(result, end-position+1)
			end = i - 1
			position = i - 1
		}
	}
	slices.Reverse(result)

	return result
}
