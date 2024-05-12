func heightChecker(heights []int) int {
	var arr = make([]int, len(heights))
	copy(arr, heights)
	counter := 0
    sort.Ints(arr)
	for i := 0; i < len(arr); i++ {
		if arr[i] != heights[i] {
			counter++
		}
	}
    return counter
}