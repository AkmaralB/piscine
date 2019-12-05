package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	first := []rune(os.Args[0])
	for _, v := range first {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}
