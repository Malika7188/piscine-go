package piscine

func IsNumeric(s string) bool {
	for _, r := range s {
		if r >= '0' && r <= '9' {
			continue
		} else if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
			return false
		} else {
			return false
		}
	}
	return true
}
