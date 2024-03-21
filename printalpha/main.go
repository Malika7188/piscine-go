package main

import "github.com/01-edu/z01"

func main() {
	for i := 'z'; i >= 'j'; i-- {
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')
}
