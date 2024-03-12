package main

import "github.com/01-edu/z01"

var str = "x = 42, y = 21"

func main() {
	for _, ch := range str {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}
