package piscine

func BasicAtoi2(s string) int {
	a := 0

	for _, v := range s {
		b := 0
		if v < '0' || v > '9' {
			return 0
		}

		for i := '1'; i <= v; i++ {
			b++
		}
		a = a*10 + b
	}
	return a
}
