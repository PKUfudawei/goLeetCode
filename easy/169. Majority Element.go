package goLeetCode

func MajorityElement(nums []int) int {
	count := map[int]int{}
	for _, element := range nums {
		if value, ok := count[element]; ok {
			count[element] = value + 1
		} else {
			count[element] = 1
		}
		if count[element] > len(nums)/2 {
			return element
		}
	}

	return 0
}
