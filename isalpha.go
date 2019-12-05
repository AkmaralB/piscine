package piscine

func IsAlpha(str string) bool {
	for _, v := range []rune(str) {
		if (v < 'A' || v > 'Z') && (v < 'a' || v > 'z') && (v < '0' || v > '9') {
			return false
		}
	}
	return true
}
