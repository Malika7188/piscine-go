package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	summary := make(map[string]int)
	var items string

	for _, i := range str {
		if i == 32 {
			summary[items] += 1
			items = ""
		} else if i != 32 {
			items += string(byte(i))
		}
	}
	summary[items] += 1
	return summary
}
