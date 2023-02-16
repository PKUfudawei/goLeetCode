package goLeetCode

func RomanToInt(s string) int {
	roman := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	if len(s) == 0 {
		return 0
	}
	result := roman[s[len(s)-1]]
	for i := len(s) - 1; i >= 1; i-- {
		if roman[s[i-1]] < roman[s[i]] {
			result -= roman[s[i-1]]
		} else {
			result += roman[s[i-1]]
		}
	}
	return result
}
