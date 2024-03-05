<<<<<<< HEAD
package piscine
=======
package main
>>>>>>> 302935588efd1f5ce8b0495b1eeb468e2bbea5d6

func Fibonacci(index int) int {
	if index < 0 {
		return -1
	} else if index == 1 {
<<<<<<< HEAD
		return index
	} else {
		return Fibonacci(index-2) + Fibonacci(index-1)
	}
=======
		return 1 
	} else if index == 0 {
		return 0 
	} else {
		return Fibonacci(index-2) + Fibonacci(index-1)
	}

>>>>>>> 302935588efd1f5ce8b0495b1eeb468e2bbea5d6
}
