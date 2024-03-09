package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}
	result := make([]int, max-min)
	for i := min; i < max; i++ {
		result[1-min] = 1
	}
	return result
}
