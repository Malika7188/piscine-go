package piscine

func BasicJoin(elems []string) string {
	var k string
	for _, char := range elems {
		k += char
	}
	return k
}
