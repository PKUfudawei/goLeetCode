package goLeetCode

func TwoSum(nums []int, target int) []int {
	num_index := map[int]int{}
	for index, value := range nums {
		gap := target - value
		if _value, ok := num_index[gap]; ok {
			_value += 0
			return []int{num_index[gap], index}
		}
		num_index[value] = index
	}
	return nil
}
