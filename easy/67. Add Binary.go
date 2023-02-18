package goLeetCode

import "strconv"

func AddBinary(a string, b string) string {
	long, short := a, b
	if len(a) < len(b) {
		long, short = short, long
	}

	residual := 0
	result := ""
	for i := 0; i < len(long); i++ {
		sum := 0
		sum += int(long[len(long)-1-i] - '0')
		if len(short)-1-i >= 0 {
			sum += int(short[len(short)-1-i] - '0')
		}
		sum += residual

		result = string(strconv.Itoa(sum%2)) + result
		residual = sum / 2
	}

	if residual > 0 {
		result = "1" + result
	}
	return result
}
