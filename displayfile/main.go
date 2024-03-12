package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	h := os.Args[1:]
	if len(h) == 0 {
		fmt.Printf("File name missing")
		fmt.Printf("\n")
	} else if len(h) > 1 {
		fmt.Printf("Too many arguments")
		fmt.Printf("\n")
	}
	ex, _ := ioutil.ReadFile(h[0])
	fmt.Printf(string(ex))
}
