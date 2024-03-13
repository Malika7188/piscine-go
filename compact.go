package piscine

func Compact(ptr *[]string) int {
	n := *ptr
	count := 0
	for _, ch := range *ptr {
		if ch != "" {
			n[count] = ch
			count++
		}
	}
	*ptr = n[0:count]
	return count
}
