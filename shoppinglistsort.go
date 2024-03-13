package piscine

func ShoppingListSort(slice []string) []string {
	for a := 0; a < len(slice)-1; a++ {
		for b := 0; b < len(slice)-a-1; b++ {
			if len(slice[b]) > len(slice[b+1]) {
				slice[b], slice[b+1] = slice[b+1], slice[b]
			}
		}
	}
	return slice
}
