package piscine

func DescendAppendRange(max, min int) []int {
	k := []int{}
	if max <= min {
		return []int{}
	}
	for i := max; i > min; i-- {
		k = append(k, i)
	}
	return k
}
