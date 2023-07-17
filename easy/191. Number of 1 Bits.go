package goLeetCode

func HammingWeight(num uint32) int {
	var result int
	for i := 0; (i < 32) && (num > 0); i++ {
		if num&1 == 1 {
			result += 1
		}
		num >>= 1
	}
	return result
}
