package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	issorted := true
	length := len(a)
	for i := 0; i < length-1; i++ {
		if f(a[i], a[i+1]) > 0 {
			issorted = false
		}
	}
	if !issorted {
		issorted = true
		for i := 0; i < length-1; i++ {
			if f(a[i], a[i+1]) < 0 {
				issorted = false
			}
		}
	}
	return issorted
}
