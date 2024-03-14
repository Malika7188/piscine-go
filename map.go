package piscine

func Map(f func(int) bool, a []int) []bool {
	n := make([]bool, len(a))
	for i := 0; i < len(a); i++ {
		if f((a[i])) {
			n[i] = true
		} else {
			n[i] = false
		}
	}
	return n
}
