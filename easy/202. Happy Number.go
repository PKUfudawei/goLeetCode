package goLeetCode

import (
	"strconv"
)

func IsHappy(n int) bool {
	visited := map[int]bool{n: true}
	for n != 1 {
		digits := strconv.Itoa(n)
		sum := 0
		for _, _rune := range digits {
			digit := int(_rune - '0')
			sum += digit * digit
		}
		if _, ok := visited[sum]; !ok {
			visited[sum] = true
		} else {
			return false
		}
		n = sum
	}
	return true
}
