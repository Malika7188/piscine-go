package piscine

func StrLen(s string) int {
	nb := 0
	for i := range []rune(s) {
		if i == i {
			nb++
		}
	}
	return nb
}
