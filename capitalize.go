package piscine

func Capitalize(s string) string {
	var capitalString string
	capitalizerFlag := true
	for _, value := range s {
		if (value >= 'a' && value <= 'z') || (value >= '0' && value <= '9') {
			if capitalizerFlag {
				if value >= 'a' && value <= 'z' {
					capitalString += string(value - 32)
				} else {
					capitalString += string(value)
				}
				capitalizerFlag = false
			} else {
				if value >= 'A' && value <= 'Z' {
					capitalString += string(value + 32)
				} else {
					capitalString += string(value)
				}
			}
		} else {
			capitalString += string(value)
			capitalizerFlag = true
		}
	}
	return capitalString
}
