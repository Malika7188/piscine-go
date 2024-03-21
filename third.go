package piscine

func ThirdTimeIsACharm(s string) string {
	if s == "" || len(s) < 3 {
		return "\n"
	}
	var result string
	for i := 2; i < len(s); i += 3 {
		result += string(s[i])
	}
	return result
}
