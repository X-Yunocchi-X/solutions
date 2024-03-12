package leetcode

import "strings"

func capitalizeTitle(title string) string {
	strs := strings.Split(title, " ")
	for i, str := range strs {
		strs[i] = strings.ToLower(str)
		if len(str) > 2 {
			strs[i] = strings.ToUpper(string(strs[i][0])) + strs[i][1:]
		}
	}
	return strings.Join(strs, " ")
}
