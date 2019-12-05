package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args[1:]
	l := 0
	for range arguments {
		l++
	}
	if l != 3 {
		return
	}
	var a, b int64
	var sf string

	as := os.Args[1]
	bs := os.Args[3]

	if IsNumeric(as) == false || IsNumeric(bs) == false {
		z01.PrintRune('0')
		z01.PrintRune('\n')
		return
	}
	a = atoi(as)
	b = atoi(bs)
	operator := os.Args[2]
	if operator == "+" {
		sf = plus(a, b)
	} else if operator == "-" {
		sf = minus(a, b)
	} else if operator == "/" {
		if b == 0 {
			sf = "No division by 0"
		} else {
			sf = div(a, b)
		}
	} else if operator == "*" {
		sf = mult(a, b)
	} else if operator == "%" {
		if b == 0 {
			sf = "No modulo by 0"
		} else {
			sf = mod(a, b)
		}
	} else {
		z01.PrintRune('0')
		z01.PrintRune('\n')
		return
	}

	ra := []rune(sf)
	for _, value := range ra {
		z01.PrintRune(value)
	}
	z01.PrintRune('\n')

}

func IsNumeric(s string) bool {
	start := 0
	l := 0
	ra := []rune(s)
	for range ra {
		l++
	}
	if ra[0] == '-' {
		start = 1
	}
	for i := start; i < l; i++ {
		if ra[i] < 48 || ra[i] > 57 {
			return false
		}
	}

	return true
}

func atoi(s string) int64 {
	var n int64
	var count int64
	n = 0
	start := 0
	l := 0
	ra := []rune(s)
	for range ra {
		l++
	}
	if ra[0] == '-' {
		start = 1
	}

	for i := start; i < l; i++ {
		count = 0
		for j := '0'; j < ra[i]; j++ {
			count++
		}
		n = n*10 + count
	}

	if ra[0] == '-' {
		n = n * -1
	}

	return n
}

func plus(a, b int64) string {
	if a > 0 && b > 0 && a+b < 0 {
		return itos(0)
	}
	if a < 0 && b < 0 && a+b > 0 {
		return itos(0)
	}
	return itos(a + b)
}
func minus(a, b int64) string {

	if itos(a) == "9223372036854775807" && b > 0 || itos(a) == "-9223372036854775807" {
		return itos(0)
	}
	if a < 0 && itos(b) == "9223372036854775807" {
		return itos(0)
	}

	return itos(a - b)
}

func mult(a, b int64) string {
	result := a * b
	if result/a != b {
		return itos(0)
	}
	return itos(a * b)
}

func div(a, b int64) string {
	return itos(a / b)
}

func mod(a, b int64) string {
	return itos(a % b)
}

func itos(n int64) string {
	s := ""
	flag := false
	if n < 0 {
		n = n * -1

		flag = true
	}

	for j := 1; j > 0; j++ {
		ost := n % 10
		switch ost {
		case 0:
			s = "0" + s
		case 1:
			s = "1" + s
		case 2:
			s = "2" + s
		case 3:
			s = "3" + s
		case 4:
			s = "4" + s
		case 5:
			s = "5" + s
		case 6:
			s = "6" + s
		case 7:
			s = "7" + s
		case 8:
			s = "8" + s
		case 9:
			s = "9" + s
		}
		n = n / 10
		if n == 0 {
			break
		}
	}
	if flag == true {
		s = "-" + s
	}
	return s

}
