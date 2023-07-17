package goLeetCode

func ContainsNearbyDuplicate(nums []int, k int) bool {
	visited := map[int]int{}
	for i, num := range nums {
		if j, ok := visited[num]; ok {
			if i-j <= k {
				return true
			}
		}
		visited[num] = i
	}
	return false
}
