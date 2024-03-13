package main

import (
	"fmt"
	"piscine"
)

func main() {
	summary := "Burger Water Carrot Coffee Water Water Chips Carrot Carrot Burger Carrot Water"
	for index, element := range piscine.ShoppingSummaryCounter(summary) {
		fmt.Println(index, "=>", element)
	}
}
