func spiralOrder(matrix [][]int) []int {
    var result []int
	y := 0
	bottom := len(matrix) - 1
	x := 0
	right := len(matrix[0]) - 1

	for {

		for i := x; i <= right; i++ {
			result = append(result, matrix[y][i])
		}
		y++
		if y > bottom {
			break
		}

		for i := y; i <= bottom; i++ {
			result = append(result, matrix[i][right])
		}
		right--
		if x > right {
			break
		}

		for i := right; i >= x; i-- {
			result = append(result, matrix[bottom][i])
		}
		bottom--
		if y > bottom {
			break
		}

		for i := bottom; i >= y; i-- {
			result = append(result, matrix[i][x])
		}
		x++
		if x > right {
			break
		}
	}
    return result
}