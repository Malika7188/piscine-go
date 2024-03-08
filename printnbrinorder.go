package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
	}
	if n < 0 {
		return
	}
	for testDigit := 0; testDigit < 10; testDigit++ {
		counter := 0
		for m := n; m > 0; m /= 10 {
			if m%10 == testDigit {
				counter++
			}
		}
		for i := 0; i < counter; i++ {
			z01.PrintRune(rune(testDigit + '0'))
		}
	}
}
