package piscine

import (
	"github.com/01-edu/z01"
)

func PrintStr(s string) {
	ary := []rune(s)
	for _, char := range ary {
		z01.PrintRune(char)
	}
}
