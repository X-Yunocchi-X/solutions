package leetcode

import (
	"strings"
)

func shortestSubstrings(arr []string) []string {
	result := make([]string, len(arr))
	for index, str := range arr {
		n := len(str)
		for i := 1; i <= n; i++ {
			if result[index] != "" {
				break
			}
			for j := 0; j < n; j++ {
				if j+i > n {
					break
				}
				temp := str[j : j+i]
				ok := true
				for index2, tempStr := range arr {
					if index != index2 && strings.Index(tempStr, temp) != -1 {
						ok = false
						break
					}
				}
				if ok {
					if result[index] == "" {
						result[index] = temp
					} else {
						result[index] = min(result[index], temp)
					}
				}
			}
		}
	}
	return result
}
