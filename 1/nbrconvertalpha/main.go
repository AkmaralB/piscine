package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args
	len := 0
	table := []string(arguments[1:])
	for index := range table {
		len = index + 1
	}
	if len > 0 {
		if table[0] == "--upper" {
			tableUpp := []string(table[1:])
			for _, num := range tableUpp {
				if Atoi(num) >= 1 && Atoi(num) <= 26 {
					z01.PrintRune(rune(Atoi(num) + 64))
				} else {
					z01.PrintRune(' ')
				}
			}
		} else {
			for _, num := range table {
				if Atoi(num) >= 1 && Atoi(num) <= 26 {
					z01.PrintRune(rune(Atoi(num) + 96))
				} else {
					z01.PrintRune(' ')
				}
			}
		}
	}
	z01.PrintRune('\n')
}

func BToInt(num byte) int {
	var runenum int
	switch num {
	case 48:
		runenum = 0
	case 49:
		runenum = 1
	case 50:
		runenum = 2
	case 51:
		runenum = 3
	case 52:
		runenum = 4
	case 53:
		runenum = 5
	case 54:
		runenum = 6
	case 55:
		runenum = 7
	case 56:
		runenum = 8
	case 57:
		runenum = 9
	}
	return runenum
}

func SLen(str string) int {
	var count int
	var a int
	for index, word := range str {
		if word == 233 {
			a = -1
		}
		count = index + 1 + a
	}
	return count
}

func BasicAtoi2(s string, ind int) int {
	lenth := SLen(s)
	var num int = 0
	str := []byte(s)
	for i := ind; i < lenth; i++ {
		if str[i] > 47 && str[i] < 58 {
			for j := i; j < lenth; j++ {
				if str[j] > 47 && str[j] < 58 {
					num = num*10 + BToInt(s[j])
				} else {
					return 0
				}
			}
			break
		} else {
			break
		}

	}
	return num
}

func Atoi(s string) int {
	str := []byte(s)
	var num int
	if SLen(s) >= 2 {

		if str[0] == 43 {
			num = BasicAtoi2(s, 1)
		} else if str[0] == 45 {
			num = (-1) * BasicAtoi2(s, 1)
		} else {
			num = BasicAtoi2(s, 0)
		}
	} else {
		num = BasicAtoi2(s, 0)
	}

	return num
}
