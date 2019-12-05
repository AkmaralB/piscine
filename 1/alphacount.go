package piscine

func AlphaCount(str string) int {
	number := 0
	for _, letters := range []rune(str) {
		if (letters >= 'A' && letters <= 'Z') || (letters >= 'a' && letters <= 'z') {
			number = number + 1
		}
	}
	return number
}
