package goLeetCode

func Merge(nums1 []int, m int, nums2 []int, n int) {
	i, j := 0, 0

	for order := i + j; j < n; order++ {
		if i < m && nums1[m-1-i] > nums2[n-1-j] {
			nums1[m+n-1-order] = nums1[m-1-i]
			i++
		} else {
			nums1[m+n-1-order] = nums2[n-1-j]
			j++
		}
	}
}
