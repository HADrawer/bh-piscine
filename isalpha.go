package piscine

func IsAlpha(s string) bool {
	result := true
	str := []rune(s)

	for _, ch := range str {
		if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || (ch >= '0' && ch <= '9') {
			result = true
		} else {
			result = false
			return result
		}
	}
	return result
}
