package piscine

func Abort(a, b, c, d, e int) int {
	n := []int{a, b, c, d, e}
	for i := 0; i < len(n); i++ {
		for j := i + 1; j < len(n); j++ {
			if n[i] > n[j] {
				n[i], n[j] = n[j], n[i]
			}
		}
	}
	return n[2]
}
