package goLeetCode

func SearchInsert(nums []int, target int) int {
	if len(nums) == 1 {
		if nums[0] >= target {
			return 0
		} else {
			return 1
		}
	}
	middle := len(nums) / 2
	if nums[middle] == target {
		return middle
	} else if nums[middle] > target {
		return SearchInsert(nums[:middle], target)
	} else {
		return middle + SearchInsert(nums[middle:], target)
	}
}
