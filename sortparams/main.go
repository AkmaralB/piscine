package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	len := 0
	for range os.Args {
		len++
	}
	for i := 1; i < len; i++ {
		for k := 1; k < len-i; k++ {
			if os.Args[k] > os.Args[k+1] {
				os.Args[k], os.Args[k+1] = os.Args[k+1], os.Args[k]
			}
		}
	}
	for i2 := 1; i2 < len; i2++ {
		for _, word := range os.Args[i2] {
			z01.PrintRune(word)
		}
		z01.PrintRune('\n')
	}
}
