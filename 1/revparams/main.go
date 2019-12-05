package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args
	l := 0

	for range arg {
		l++
	}

	for i := l - 1; i >= 1; i-- {
		arg2 := arg[i]
		for _, v := range arg2 {
			z01.PrintRune(v)
		}
		z01.PrintRune('\n')
	}
}
