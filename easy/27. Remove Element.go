package goLeetCode

func RemoveElement(nums []int, val int) int {
	if len(nums) < 1 {
		return 0
	}

	unmatch := 0
	for _, num := range nums {
		if num != val {
			nums[unmatch] = num
			unmatch++
		}
	}
	return unmatch
}
