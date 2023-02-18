package goLeetCode

import "strings"

func LengthOfLastWord(s string) int {
	i := len(s) - 1
	for ; i >= 0 && s[i] == ' '; i-- {
	}
	end := i
	for ; i >= 0 && s[i] != ' '; i-- {
	}
	start := i
	return end - start
}

func LengthOfLastWord2(s string) int {
	str := strings.Fields(s) //tokenize
	return len(str[len(str)-1])
}
