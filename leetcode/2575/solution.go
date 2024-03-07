package leetcode

func divisibilityArray(word string, m int) []int {
	result := make([]int, len(word))
	mod := 0
	for i, char := range word {
		mod = (mod*10 + int(char-'0')) % m
		if mod == 0 {
			result[i] = 1
		}
	}
	return result
}
