package piscine

func IsLower(s string) bool {
	for _, r := range s {
		if r >= 'a' && r <= 'z' {
			continue
		} else if r >= 'A' && r <= 'Z' {
			return false
		} else {
			return false
		}
	}
	return true

}
