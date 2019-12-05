package piscine

func NRune(s string, n int) rune {
	runes := []rune(s)
	for index, value := range runes {
		if index+1 == n {
			return value
		}
	}
	return 0
}
