func reverseString(s []byte)  {
	var swap byte
	for start, end := 0, len(s)-1; start < end; start, end = start+1, end-1 {
		swap = s[start]
		s[start] = s[end]
		s[end] = swap
	}
}