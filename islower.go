package piscine

func IsLower(s string) bool {
	result := true
	str := []rune(s)

	for _, ch := range str {
		if ch >= 'a' && ch <= 'z' {
			result = true
		} else {
			result = false
			return result
		}
	}
	return result
}
