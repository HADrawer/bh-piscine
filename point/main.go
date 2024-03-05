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
	z01.PrintRune(52)
	z01.PrintRune(50)
	for _, r := range b {
		z01.PrintRune(r)
	}
	z01.PrintRune(50)
	z01.PrintRune(49)
	z01.PrintRune('\n')
}
