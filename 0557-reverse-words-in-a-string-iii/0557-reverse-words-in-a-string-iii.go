func reverseWords(s string) string {
    var reverse string
	pointer := 0
	counter := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			pointer = i-1
			for counter <= pointer {
				reverse += string(s[pointer])
				pointer--
			}
            reverse+=" "
			counter = i + 1
		} else if i == len(s)-1 {
			pointer = i
			for counter <= pointer {
				reverse += string(s[pointer])
				pointer--
			}
		}
	}
    return reverse
}