package piscine

func ToUpper(s string) string {
	str := []rune(s)
	for i := 0; i <= len(str)-1; i++ {
		if str[i] >= 97 && str[i] <= 122 {
			str[i] = str[i] - 32
		}
	}
	return string(str)
}
