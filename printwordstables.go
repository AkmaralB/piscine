package piscine

import (
	"github.com/01-edu/z01"
)

func PrintWordsTables(table []string) {
	for _, v := range table {
		for _, v2 := range v {
			z01.PrintRune(v2)
		}
		z01.PrintRune('\n')
	}
}
