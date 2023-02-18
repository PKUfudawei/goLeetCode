package goLeetCode

func MySqrt(x int) int {
	answer := x
	for answer*answer > x {
		answer = (answer + x/answer) / 2
	}
	return answer
	/*compare := (float32(answer) + 0.50) * (float32(answer) + 0.50)
	if compare > float32(x) {
		return answer
	} else {
		return answer + 1
	}*/
}
