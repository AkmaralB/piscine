package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
	}
	PrintNumber(n)
}

func PrintNumber(number int) {
	i := '0'
	if number == 0 {
		z01.PrintRune('0')
		return
	}

	for a := 1; a <= number%10; a++ {
		i++
	}

	for b := -1; b >= number%10; b-- {
		i++
	}

	if number/10 != 0 {
		PrintNumber(number / 10)
	}
	z01.PrintRune(i)
	return
}
