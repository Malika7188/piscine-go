package piscine

func LoafOfBread(str string) string {
	if str == "" || len(str) < 5 {
		return "\n"
	}
	result := ""
	for i := 0; i < len(str); i++ {
		if i != 0 {
			if (i+1)%5 == 2 {
				result += string(str[i])
			}
		}
	}
	result += "\n"
	return result
}
