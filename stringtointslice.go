package piscine

func StringToIntSlice(str string) []int {
	var result []int
	for _, ch := range str {
		result = append(result, int(ch))
	}
	return result
}
