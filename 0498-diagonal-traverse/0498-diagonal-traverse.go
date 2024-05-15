func findDiagonalOrder(mat [][]int) []int {
    i, j := 0, 0
	x := len(mat) - 1
	y := len(mat[0]) - 1
	var result []int
	up := true
	for {
		if i+j > x+y {
			break
		}
		result = append(result, mat[i][j])

		if up && j >= y {
			up = false
			i++
		} else if up && i <= 0 {
			up = false
			j++
		} else if !up && i >= x {
			up = true
			j++
		} else if !up && j <= 0 {
			up = true
			i++
		} else if up {
			i--
			j++
		} else {
			i++
			j--
		}
	}
	return result
}