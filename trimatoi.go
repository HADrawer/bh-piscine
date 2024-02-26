package piscine

func TrimAtoi(s string) int {
	num := 0
	minus := false
	str := []rune(s)
	for _, ch := range str {
		lastdigit := 0
		if ch <= '9' && ch >= '0' {
			for i := '1'; i <= ch; i++ {
				lastdigit++
			}
			num = num*10 + lastdigit
		} else if num == 0 && ch == '-' {
			minus = true
			continue
		} else {
			continue
		}
	}
	if minus == true {
		num = num * -1
	}
	return num
}
