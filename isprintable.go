package piscine

func IsPrintable(s string) bool {
	result := true
	str := []rune(s)

	for _, ch := range str {
		if ch >= 32 && ch <= 127 {
			result = true
		} else {
			result = false
			return result
		}
	}
	return result
}
