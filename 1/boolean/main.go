package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printStr(str string) {

	for _, v := range str {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}

func isEven(nbr int) bool {
	if nbr%2 == 1 {
		return false
	}
	return true

}

func main() {
	if isEven(lengthOfArg(os.Args[1:])) {
		printStr("I have an even number of arguments")
	} else {
		printStr("I have an odd number of arguments")
	}
}
func lengthOfArg(arr []string) int {
	lens := 0
	for range arr {
		lens++
	}
	return lens
}
