package piscine

func ReverseMenuIndex(menu []string) []string {
	n := len(menu)
	for i := 0; i < n/2; i++ {
		menu[i], menu[n-i-1] = menu[n-i-1], menu[i]
	}
	return menu
}
