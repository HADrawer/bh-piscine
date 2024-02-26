package piscine

func IsNumeric(s string) bool {
	result := true
	str := []rune(s)

	for _, ch := range str {
		if ch >= '0' && ch <= '9' {
			result = true
		} else {
			result = false
			return result
		}
	}
	return result
}
