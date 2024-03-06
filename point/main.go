package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}
	setPoint(points)
	a := "x = "
	b := ", y = "
	for _, r := range a {
		z01.PrintRune(r)
	}
	printNumber(points.x)
	for _, r := range b {
		z01.PrintRune(r)
	}
	printNumber(points.y)
	z01.PrintRune('\n')
}
func printDigit(n int) {
	if n/10 != 0 {
		printDigit(n / 10)
	}
	z01.PrintRune(rune(n%10) + '0')
}
func printNumber(n int) {
	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}
	printDigit(n)
}
