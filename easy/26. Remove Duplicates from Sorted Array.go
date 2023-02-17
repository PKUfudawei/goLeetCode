package goLeetCode

func RemoveDuplicates(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	unique := 0
	for i, num := range nums {
		if num != nums[unique] {
			unique += 1
			nums[unique] = nums[i]
		}
	}
	return unique + 1
}
