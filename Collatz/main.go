package main

import (
	"fmt"
	"piscine"
)

func main() {
	steps := piscine.CollatzCountdown(12)
	fmt.Println(steps)
}
