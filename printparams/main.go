package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args
	for i, vArg := range arg {
		if i != 0 {
			for _, vWord := range vArg {
				z01.PrintRune(vWord)
			}
			z01.PrintRune('\n')
		}
	}
}
