package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	argument := os.Args[1:]
	for i := 0; i < len(argument)-1; i++ {
		for j := 0; j < len(argument)-i-1; j++ {
			if argument[j] > argument[j+1] {
				argument[j], argument[j+1] = argument[j+1], argument[j]
			}
		}
	}
	for _, str := range argument {
		for _, cha := range str {
			z01.PrintRune(cha)
		}
		z01.PrintRune('\n')
	}
}
