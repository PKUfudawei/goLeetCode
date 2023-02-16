package goLeetCode

import (
	"strconv"
)

func IsPalindrome(x int) bool {
	runes := []rune(strconv.Itoa(x))
	length := len(runes)

	for index, rune := range runes {
		if rune != runes[length-1-index] {
			return false
		}
	}
	/*for i := 0; i <= length/2; i++ {
		if s[i] != s[length-1-i] {
			return false
		}
	}*/
	return true
}
