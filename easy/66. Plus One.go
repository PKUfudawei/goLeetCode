package goLeetCode

func PlusOne(digits []int) []int {
	digits[len(digits)-1] += 1

	for i := len(digits) - 1; i > 0; i-- {
		if digits[i] < 10 {
			break
		} else {
			digits[i] = 0
			digits[i-1] += 1
		}
	}

	if digits[0] > 9 {
		digits[0] = 0
		digits = append([]int{1}, digits...)
	}
	return digits
}
