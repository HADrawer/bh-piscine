package piscine

func StrRev(s string) string {
	ary := []rune(s)
result := ""
	for i := len(ary)-1; i >= 0 ; i--{
	result += string(ary[i])
	}
	return result
}