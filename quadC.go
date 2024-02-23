package piscine

import "github.com/01-edu/z01"

func QuadC(x, y int) {
	if x >= 0 && y >= 0 {
		for i := 0; i <= y-1; i++ {
			for k := 0; k <= x-1; k++ {
				if i == 0 && k == 0 {
					z01.PrintRune('A')
				}
				if i == y-1 && k == 0 {
					z01.PrintRune('C')
				}
				if (i == 0 || i == y-1) && (k > 0 && k < x-1) {
					z01.PrintRune('B')
				}
				if i > 0 && i < y-1 && k == 0 {
					z01.PrintRune('B')
				}
				if i > 0 && i < y-1 && k > 0 && k < x-1 {
					z01.PrintRune(' ')
				}
				if i > 0 && i < y-1 && k == x-1 {
					z01.PrintRune('B')
				}

				if i == 0 && k == x-1 && x-1 != 1 {
					z01.PrintRune('A')
				}
				if i == y-1 && k == x-1 && x-1 != 1 {
					z01.PrintRune('C')
				}
			}
			z01.PrintRune('\n')
		}
	}
}
