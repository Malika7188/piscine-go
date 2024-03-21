package main

import (
	"os"
	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 2 {
		arg := []rune(os.Args[1])
		for i, ch := range arg {
			if ch >= 'a' && ch <= 'z' {
				arg[i] = 'z' - ch + 'a'
			} else if ch >= 'A' && ch <= 'Z' {
				arg[i] = 'Z' - ch + 'A'
			}
			z01.PrintRune(rune(arg[i]))
		}
		
	}

	z01.PrintRune('\n')
}
