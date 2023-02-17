package goLeetCode

func IsValid(s string) bool {
	stack := []rune{}
	right := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	left := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, rune_ := range s {
		if value, ok := right[rune_]; ok {
			stack = append(stack, value)
		} else if _, ok := left[rune_]; ok {
			if len(stack) > 0 && stack[len(stack)-1] == rune_ {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}

	return len(stack) == 0
}
