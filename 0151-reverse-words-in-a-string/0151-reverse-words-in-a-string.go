func reverseWords(s string) string {
	arr := strings.Fields(s)
	start := 0
	end := len(arr) - 1
	for start < end {
		arr[start], arr[end] = arr[end], arr[start]
		start++
		end--
	}
	words := arr[0]
	for i := 1; i < len(arr); i++ {
		words += " " + arr[i]
	}
    return words
}