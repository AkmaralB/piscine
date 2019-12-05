package piscine

func ToUpper(s string) string {
	runes := []rune(s)
	for index, value := range runes {
		i := index
		if value >= 97 && value <= 122 {
			runes[i] = value - 32
		}
	}
	return string(runes)
}
