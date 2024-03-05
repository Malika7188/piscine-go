package main

func Fibonacci(index int) int {
	if index < 0 {
		return -1
	} else if index == 1 {
		return 1 
	} else if index == 0 {
		return 0 
	} else {
		return Fibonacci(index-2) + Fibonacci(index-1)
	}

}
