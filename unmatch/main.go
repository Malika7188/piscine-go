package main

import (
	"fmt"
	"piscine"
)

func main() {
	a := []int{1, 2, 3, 1, 2, 3, 4}
	unmatch := piscine.Unmatch(a)
	fmt.Println(unmatch)
}
