package goLeetCode

func TitleToNumber(columnTitle string) int {
	result := 0
	for _, element := range columnTitle {
		result = result*26 + int(element-'A'+1)
	}
	return result
}
