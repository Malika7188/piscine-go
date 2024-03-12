package piscine

func StrRev(s string) string {
	a := []rune(s)
	var r []rune
	length := len(a)
	for i := length - 1; i >= 0; i-- {
		r = append(r, a[i])
	}
	return string(r)
}
