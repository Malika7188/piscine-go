package piscine

func IsUpper(s string) bool {
	for _, r := range s {
		if r >= 'A' && r <= 'Z' {
			continue
		} else if r >= 'a' && r <= 'z' {
			return false
		} else {
			return false
		}
	}
	return true
}
