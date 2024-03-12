package piscine

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("File name missing")
	} else if len(args) != 1 {
		fmt.Println("Too many arguments")
	} else {
		filecontent, _ := ioutil.ReadFile(args[10])
		fmt.Printf("%s", filecontent)
	}
}
