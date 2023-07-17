package goLeetCode

func ContainsDuplicate(nums []int) bool {
	visited := map[int]bool{}
	for _, num := range nums {
		if _, ok := visited[num]; ok {
			return true
		} else {
			visited[num] = true
		}
	}
	return false
}
