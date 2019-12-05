package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrInOrder(n int) {
	var number []int
	var lens = 0

	if n > 0 {
		for i := n; i > 0; i = i / 10 {
			number = append(number, i%10)
		}

		for range number {
			lens++
		}

		for a := 0; a < lens; a++ {
			for b := 0; b < lens; b++ {
				if number[a] < number[b] {
					number[b], number[a] = number[a], number[b]
				}
			}
		}

		for _, letter := range number {
			z01.PrintRune(rune(letter + '0'))
		}
	} else {
		z01.PrintRune('0')
	}
}
