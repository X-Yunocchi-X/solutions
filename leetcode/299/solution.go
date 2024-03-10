package leetcode

import "fmt"

func getHint(secret string, guess string) string {
	bulls, cows := 0, 0
	sCnt, gCnt := [10]int{}, [10]int{}
	for i := range secret {
		if secret[i] == guess[i] {
			bulls++
		}
	}
	for i := range secret {
		if secret[i] == guess[i] {
			continue
		}
		sCnt[secret[i]-'0']++
		gCnt[guess[i]-'0']++
	}
	for i := range sCnt[:] {
		cows += min(sCnt[i], gCnt[i])
	}
	return fmt.Sprintf("%dA%dB", bulls, cows)
}
