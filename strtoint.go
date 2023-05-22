package proj

func StringToInt(str string) int {
	result := 0
	for _, c := range str {
		if c >= '0' && c <= '9' {
			result = result*10 + int(c-'0')
		} else {
			continue
		}
	}
	return result
}
