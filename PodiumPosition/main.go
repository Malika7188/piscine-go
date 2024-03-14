package main

import (
	"fmt"
	"piscine"
)

func main() {

	p4 := []string{"4th Place"}
	p3 := []string{"3rd Place"}
	p2 := []string{"2nd Place"}
	p1 := []string{"1st Place"}

	position := [][]string{p4, p3, p2, p1}
	fmt.Println(piscine.PodiumPosition(position))
}
