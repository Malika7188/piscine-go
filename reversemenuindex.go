package piscine

func ReverseMenuIndex(menu []string) []string {
	arg := make([]string, len(menu))

	for i := len(menu) - 1; i >= 0; i-- {
		arg[len(menu)-1-i] = menu[i]
	}
	return arg
}

