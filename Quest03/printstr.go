package piscine

import "github.com/01-edu/z01"

func PrintStr(str string) {
	Byte := []byte(str)

	for _, letter := range Byte {
		z01.PrintRune(rune(letter))
	}
}
