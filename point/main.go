package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func PR(r rune) {
	z01.PrintRune(r)
}

func PrintStr(s string) {
	for _, v := range s {
		PR(v)
	}
}

func check(r int) {
	c := '0'
	if r == 0 {
		PR(c)
	}
	for i := 1; i <= r%10; i++ {
		c++
	}
	for i := -1; i >= r%10; i-- {
		c++
	}
	if r/10 != 0 {
		check(r / 10)
	}
	PR(c)
}

func main() {
	points := &point{}
	setPoint(points)
	PrintStr("x = ")
	check(points.x)
	PrintStr(", y = ")
	check(points.y)
	PrintStr("\n")
}
