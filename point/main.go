package main

import (
	"fmt"

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

func main() {
	points := &point{}
	setPoint(points)
	z01.PrintRune('x')
	fmt.Printf(" = %d, y = %d\n", points.x, points.y)
}
