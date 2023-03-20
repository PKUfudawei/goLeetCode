package goLeetCode

func GetRow(rowIndex int) []int {
	values := [][]int{}
	for row := 0; row <= rowIndex; row++ {
		rowValues := []int{1}
		for column := 1; column < row; column++ {
			rowValues = append(rowValues, values[row-1][column-1]+values[row-1][column])
		}
		if row > 0 {
			rowValues = append(rowValues, 1)
		}
		values = append(values, rowValues)
	}
	return values[len(values)-1]
}
