package main

import "github.com/01-edu/z01"

func main() {
	PrintNbr(-123)
	PrintNbr(0)
	PrintNbr(123)
}

func PrintNbr(n int) {
	nbr := string(n)
	for _, v := range nbr {
		z01.PrintRune(rune(v))
	}
}
