package goLeetCode

import (
	"strings"
)

func cleanString(s string) string {
	result := []rune{}
	for _, char := range strings.ToLower(s) {
		if ('a' <= char && char <= 'z') || ('0' <= char && char <= '9') {
			result = append(result, char)
		}
	}
	return string(result)
}

func IsPalindromeString(s string) bool {
	s = cleanString(s)
	for i := 0; i < len(s); i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
