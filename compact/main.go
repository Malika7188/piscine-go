package main

import (
	"fmt"
	"piscine"
)

const N = 6

func main() {
	a := make([]string, N)
	a[0] = "a"
	a[2] = "b"
	a[4] = "c"

	for _, v := range a {
		fmt.Println(v)
	}

	fmt.Println("Size after compacting:", piscine.Compact(&a))

	for _, v := range a {
		fmt.Println(v)
	}
}
