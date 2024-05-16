func addBinary(a string, b string) string {
    var summa string
	carry := 0
	i := len(a) - 1
	j := len(b) - 1
	for i >= 0 || j >= 0 {
		sum := carry
		if i >= 0 {
			sum += int(a[i] - '0')
			i--
		}
		if j >= 0 {
			sum += int(b[j] - '0')
			j--
		}
		summa = fmt.Sprint(sum%2) + summa
		carry = sum / 2
	}
	if carry > 0 {
		summa = "1" + summa
	}
	return summa
}