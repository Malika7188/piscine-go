package piscine

func CountChar(str string, char byte) int {
	if str == "" {
		return 0
	}
	count := 0
	for i := 0; i < len(str); i++ {
		if str[i] == char {
			count++
		}
	}
	return count
}
