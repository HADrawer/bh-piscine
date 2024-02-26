package piscine

func Pow(a, b int) int {
	p := 1
	for b > 0 {
		if b&1 != 0 {
			p *= a
		}
		b >>= 1
		a *= a
	}
	return p
}
func BasicAtoi(s string) int {

	ary := []rune(s)
	result := 0

	for i := 0; i <= len(ary)-1; i++ {
		if ary[i] > 0 {
			result += int(ary[i])
			result = result * 10
		}
	}

	return result
}
