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
	str := "x = 42, y = 21"
	for _, v := range str {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}
