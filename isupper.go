package piscine

func IsUpper(s string) bool {
	result := true
	str := []rune(s)

	for _, ch := range str {
		if ch >= 'A' && ch <= 'Z' {
			result = true
		} else {
			result = false
		}
	}
	return result
}
