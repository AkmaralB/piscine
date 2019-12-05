package piscine

func IsNumeric(str string) bool {
	for _, v := range []rune(str) {
		if v < '0' || v > '9' {
			return false
		}
	}
	return true
}
