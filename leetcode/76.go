package leetcode

func minWindow(s string, t string) string {
	if len(t) > len(s) {
		return ""
	}

	res := ""
	countT := [256]int{}
	countS := [256]int{}
	zero := 0

	for _, char := range t {
		if countT[char] == 0 {
			zero++
		}
		countT[char]++
	}

	start := 0

	for end, char := range s {
		countS[char]++
		if countS[char] == countT[char] {
			zero--
		}
		for zero == 0 {
			temp := s[start : end+1]
			if res == "" || len(temp) < len(res) {
				res = temp
			}
			char := s[start]
			if countS[char] == countT[char] && countT[char] != 0 {
				zero++
			}
			countS[char]--
			start++
		}
	}
	return res
}
