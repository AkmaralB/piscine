package piscine

func StrLen(str string) int {
	nb := 0
	for i := range []rune(str) {
		if i == i {
			nb++
		}
	}
	return nb
}
