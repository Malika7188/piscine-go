package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	argument := os.Args[1:]
	for _, val := range argument {
		for _, char := range val {
			z01.PrintRune(char)
		}
		z01.PrintRune('\n')
	}
}
