package piscine

func IsUpper(str string) bool {
	for _, v := range []rune(str) {
		if v < 'A' || v > 'Z' {
			return false
		}
	}
	return true
}
