package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for i := '0'; i <= '7'; i++ {
		for j := i + 1; i <= '8'; j++ {
			for k := j + 1; k <= '9'; k++ {
				if i < j {
					if j < k {
						z01.PrintRune(i)
						z01.PrintRune(j)
						z01.PrintRune(k)
						if i == '7' && j == '8' && k == '9' {
							z01.PrintRune('\n')
						} else {
							z01.PrintRune(',')
							z01.PrintRune(' ')
						}
					}
				}
			}
		}
	}
}
