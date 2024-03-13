package piscine

func Unmatch(a []int) int {
	s := make(map[int]bool)
	for _, v := range a {
		if !s[v] {
			s[v] = true
			return v
		}
	}
	return -1
}
