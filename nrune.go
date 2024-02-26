package piscine

func NRune(s string, n int) rune {
result := []rune(s)

if n > len(result)  || n < 0{
	return 0
}

return result[n-1]
}