//Write a function ThirdTimeIsACharm() that takes a string as an argument and returns another string with every third character.
//Return the output followed by a newline \n.
//If the string is empty, return a newline \n.
//If there is no third character, return a newline \n.

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
