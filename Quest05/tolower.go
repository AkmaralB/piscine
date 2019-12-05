package piscine

func ToLower(s string) string {
	runes := []rune(s)
	for index, value := range runes {
		if value >= 'A' && value <= 'Z' {
			runes[index] = value + 32
		}
	}
	return string(runes)
}
