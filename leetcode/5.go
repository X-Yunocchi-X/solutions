package leetcode

func longestPalindrome(s string) string {
	if len(s) == 1 {
		return s
	}
	res := ""

	for i := range s {
		for j := 0; j <= min(i, len(s)-i-1); j++ {
			if s[i-j] != s[i+j] {
				break
			}
			length := 1 + 2*j
			if length > len(res) {
				res = s[i-j : i+j+1]
			}
		}
		for j := 0; j <= min(i, len(s)-i-2); j++ {
			if s[i-j] != s[i+1+j] {
				break
			}
			length := 2 + 2*j
			if length > len(res) {
				res = s[i-j : i+j+2]
			}
		}
	}

	return res
}
