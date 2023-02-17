package goLeetCode

func LongestCommonPrefix(strs []string) string {
	common_prefix := strs[0]

	for _, str := range strs {
		for j := 0; j < len(common_prefix); j++ {
			if j+1 > len(str) || str[j] != common_prefix[j] {
				common_prefix = common_prefix[:j]
				break
			}
		}
	}

	return common_prefix
}
