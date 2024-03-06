package piscine

func Sqrt(nb int) int {
	if nb == 0 {
		return 0
	}
	if nb == 1 {
		return 1
	}
	h := nb / 2

	for a := 1; a < 10; a++ {
		if h != 0 {
			h = (h + nb/h) / 2
		} else {
			break
		}
	}
	if h*h != nb {
		return 0
	} else {
		return h
	}
}
