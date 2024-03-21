package piscine

import "fmt"

func BuzZinga(nb int) {
	if nb < 0 {
		fmt.Println("*")
	}
	if nb == 0 {
		fmt.Println("BuzZinga")
	}
	for i := 1; i <= nb; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("BuzZinga")
		} else if i%3 == 0 {
			fmt.Println("Buz")
		} else if i%5 == 0 {
			fmt.Println("Zinga")
		} else {
			fmt.Println("*")
		}
	}
}
