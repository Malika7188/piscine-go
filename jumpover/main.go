package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Print(piscine.JumpOver("1010101010"))
	fmt.Print(piscine.JumpOver(""))
	fmt.Print(piscine.JumpOver("t w e l v e"))
	fmt.Print(piscine.JumpOver("12"))
}
