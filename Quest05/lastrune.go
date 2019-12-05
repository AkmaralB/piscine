package piscine

func LastRune(s string) rune {
	runes := []rune(s)
	lens := 0
	for range runes {
		lens++
	}
	for index, value := range runes {
		if index == lens-1 {
			return value
		}
	}
	return 0
}
