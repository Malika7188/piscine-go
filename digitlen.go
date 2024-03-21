//Write a function DigitLen() that takes two integers as arguments and returns the times the first int can be divided by the second until it reaches zero.
//The second int must be between 2 and 36. If not, the function returns -1.
//If the first int is negative, reverse the sign and count the digits.
package piscine

func DigitLen(n, base int) int {
	if base < 2 || base > 36 {
		return -1
	}
	if n < 0 {
		n = -n
	}
	count := 0
	for n > 0 {
		n /= base
		count++
	}
	return count
}
