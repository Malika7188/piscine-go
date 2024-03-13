package piscine

func Unmatch(a []int) int {
	for _, res := range a {
		test := 0
		for _, el := range a {
			if el == res {
				test++
			}
		}
		if test == 1 || test%2 == 1 {
			return res
		}
	}
	return -1
}
