package piscine

func JumpOver(str string) string {
	if str == "" || len(str) < 3 {
		return "\n"
	}
	result := ""
	for i := 0; i < len(str); i++ {
		if i != 0 {
			if (i+1)%3 == 0 {
				result += string(str[i])
			}
		}
	}
	result += "\n"
	return result
}
