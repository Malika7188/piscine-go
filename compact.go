package piscine

func Compact(ptr *[]string) int {
	n := *ptr
	count := 0
	for _, ch := range *ptr {
		if ch != "" {
			n = append(n, ch)
			count++
		}
	}
	*ptr = n
	return count
}
