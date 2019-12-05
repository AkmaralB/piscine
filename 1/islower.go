package piscine

func IsLower(str string) bool {
	for _, v := range []rune(str) {
		if v < 'a' || v > 'z' {
			return false
		}
	}
	return true
}
