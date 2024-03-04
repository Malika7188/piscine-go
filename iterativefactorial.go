package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 || nb > 20 {
		return 0
	} else if nb == 0 {
		return 1
	} else {
		k := nb - 1
		for i := k; i >= 1; i-- {
			nb *= i
		}
	}
	return nb

}
