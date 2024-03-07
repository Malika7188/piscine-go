package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for i := '0'; i <= '6'; i++ {
		for j := i + 1; i <= '7'; j++ {
			for k := j + 1; i <= '8'; k++ {
				z01.PrintRune(i)
				z01.PrintRune(j)
				z01.PrintRune(k)
				if i == 6 || j == 7 || k == 8 {
					return
				}
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}

}
