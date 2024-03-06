package piscine

func LastRune(s string) rune {
	last := []rune(s)
	return last[len(s)-1]
}
