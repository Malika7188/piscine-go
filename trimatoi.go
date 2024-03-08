package piscine

func TrimAtoi(s string) int {
	num := 0
	symbol := 1
	for _, char := range s {
		if char >= '0' && char <= '9' {
			num = num*10 + int(char-'0')
		} else if char == '-' && num == 0 {
			symbol = -1
		} else if num > 0 {
			continue
		}
	}
	return symbol * num
}
