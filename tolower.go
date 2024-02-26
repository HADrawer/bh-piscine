package piscine

func ToLower(s string) string {
	str := []rune(s)
	for i := 0; i <= len(str)-1; i++ {
		if str[i] >= 65 && str[i] <= 90 {
			str[i] = str[i] + 32
		}
	}
	return string(str)
}
