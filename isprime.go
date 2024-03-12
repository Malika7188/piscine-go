package piscine

func IsPrime(nb int) bool {
	if nb <= 1 {
		return false
	}
	if nb <= 3 {
		return true
	}
	if nb%2 == 0 || nb%3 == 0 {
		return false
	}
	k := 5
	for k*k <= nb {
		if nb%k == 0 || nb%(k+2) == 0 {
			return false
		}
		k += 6
	}
	return true
}
