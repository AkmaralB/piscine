package piscine

func BasicAtoi(s string) int {
	a := 0
	for _, i := range s {
		b := 0
		for k := '1'; k <= i; k++ {
			b++
		}
		a = a*10 + b
	}
	return a
}
