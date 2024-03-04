package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	name := os.Args[1:]
	length := len(name) - 1
	for length >= 0 {

		for _, str := range name[length] {
			z01.PrintRune(str)
		}
		length--
		z01.PrintRune('\n')
	}
}
