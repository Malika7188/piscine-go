package main

import "os"

func main() {
	argument := os.Args[0][2:]
	for _, val := range argument {
		z01.PrintRune(val)
	}
	z01.PrintRune('\n')
}
