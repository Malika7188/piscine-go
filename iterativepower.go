package piscine

func IterativePower(nb int, power int) int {
	if power < 0 {
		return 0
	} else {
		m := 1
		for a := 1; a <= power; a++ {
			m = m * nb
		}
		return m
	}
}
