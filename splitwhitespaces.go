package piscine

func SplitWhiteSpaces(s string) []string {
	var words []string
	var word string
	for _, char := range s {
		if char == ' ' || char == '\t' || char == '\n' {
			if len(word) > 0 {
				words = append(words, word)
				word = ""
			}
		} else {
			word += string(char)
		}
	}
	if len(word) > 0 {
		words = append(words, word)
	}
	return words
}
