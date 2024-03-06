package piscine

func NRune(s string, n int) rune {
	m := []rune(s)
	if n <= 0 || n > len(m) {
		return 0
	}
	return m[n-1]
}
