package main

import "github.com/01-edu/z01"

type point struct {
	x string
	y string
}

func setPoint(ptr *point) {
	ptr.x = "42"
	ptr.y = "21"
}

func main() {
	points := &point{}

	setPoint(points)
	arr := "x = " + points.x + ", y = " + points.y
	for _, v := range arr {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')

}
