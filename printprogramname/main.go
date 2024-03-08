package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	argument := os.Args[0][2:]
	for _, val := range argument {
		z01.PrintRune(val)
	}
	z01.PrintRune('\n')
}
