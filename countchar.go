//write a function that takes a string and xter as an argument and return the number of times the xter appears in the string
//if xter is not in the string return 0
//if the string is empty return 0

package piscine

func CountChar(str string, char byte) int {
	if str == "" {
		return 0
	}
	count := 0
	for i := 0; i < len(str); i++ {
		if str[i] == char {
			count++
		}
	}
	return count
}
