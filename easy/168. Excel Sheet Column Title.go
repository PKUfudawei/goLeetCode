package goLeetCode

func ConvertToTitle(columnNumber int) string {
	result := []byte{}
	for columnNumber != 0 {
		result = append(result, byte((columnNumber-1)%26)+'A')
		columnNumber = (columnNumber - 1) / 26
	}
	for i := 0; i < len(result)/2; i++ {
		result[i], result[len(result)-1-i] = result[len(result)-1-i], result[i]
	}
	return string(result)
}
