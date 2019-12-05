package piscine

func Capitalize(s string) string {
	runes := []rune(s)
	for index, value := range s {
		if value >= 'A' && value <= 'Z' {
			runes[index] = value + 32
		}
	}
	for i, v := range runes {
		if v >= 'a' && v <= 'z' && i != 0 {
			if (runes[i-1] < 'A' || runes[i-1] > 'Z') && (runes[i-1] < 'a' || runes[i-1] > 'z') && (runes[i-1] < '0' || runes[i-1] > '9') {
				runes[i] = v - 32
			}
		} else if i == 0 && v >= 'a' && v <= 'z' {
			runes[i] = v - 32
		}
	}
	return string(runes)
}
