package leetcode

import (
	"strconv"
	"strings"
	"unicode"
)

func decodeString(s string) string {
	stack := []any{} // allow for rune, string and int
	analyse := func(str string) (any, string) {
		if str[0] == '[' || str[0] == ']' {
			return rune(str[0]), str[1:]
		}
		isDigit := str[0] <= '9' && str[0] >= '0'
		end := 0
		for ; end < len(str); end++ {
			if (isDigit && unicode.IsDigit(rune(str[end]))) || (!isDigit && unicode.IsLetter(rune(str[end]))) {
				continue
			}
			break
		}
		if isDigit {
			res, _ := strconv.Atoi(str[:end])
			return res, str[end:]
		}
		return str[:end], str[end:]
	}
	for len(s) > 0 {
		char, other := analyse(s)
		stack = append(stack, char)
		if v, ok := stack[len(stack)-1].(rune); ok && v == ']' {
			str := stack[len(stack)-2].(string)
			num := stack[len(stack)-4].(int)
			stack = stack[:len(stack)-4]
			stack = append(stack, strings.Repeat(str, num))
		}
		if len(stack) > 1 {
			v1, ok1 := stack[len(stack)-1].(string)
			v2, ok2 := stack[len(stack)-2].(string)
			if ok1 && ok2 {
				stack = stack[:len(stack)-1]
				stack[len(stack)-1] = v2 + v1
			}
		}
		s = other
	}
	return stack[0].(string)
}
