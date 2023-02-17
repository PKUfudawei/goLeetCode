package goLeetCode

func LongestCommonPrefix(strs []string) string {
	common_prefix := strs[0]

	for i := 0; i < len(common_prefix); i++ {
		for _, str := range strs {
			if i+1 > len(str) || str[i] != common_prefix[i] {
				common_prefix = common_prefix[:i]
				break
			}
		}
	}

	return common_prefix
}
