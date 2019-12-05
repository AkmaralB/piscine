package piscine

func IsPrintable(str string) bool {
	for _, v := range []rune(str) {
		if !(v >= 32 && v <= 127) {
			return false
		}
	}
	return true
}
