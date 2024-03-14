package piscine

func LoafOfBread(str string) string {
	if str == "" {
		return "\n"
	}
	if len(str) < 5 {
		return "Invalid Output\n"
	}
	q := ""
	counter := 0
	for i, v := range str {
		if v != ' ' && counter != 5 {
			q += string(v)
			counter++
		} else if counter == 5 {
			q += " "
			counter = 0
		}
		if i == len(str)-1 && len(q) > 0 && q[len(q)-1] == ' ' {
			q = q[:len(q)-1]
		}
	}
	q += "\n"
	return q
}
