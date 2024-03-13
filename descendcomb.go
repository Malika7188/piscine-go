package piscine

import "github.com/01-edu/z01"

func printer(s int) {
	unit := s % 10
	test := s / 10

	z01.PrintRune(rune(unit) + '0')
	z01.PrintRune(rune(test) + '0')
}

func DescendComb() {
	for i := 99; i > 0; i-- {
		for j := i - 1; j >= 0; j-- {
			printer(i)
			z01.PrintRune(' ')
			printer(j)
			if !(i == 1 && j == 0) {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
}
