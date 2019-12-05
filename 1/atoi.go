package piscine

func Atoi(s string) int {
	sign := 1
	a := 0
	count := 0
	for i, v := range s {
		if v == '-' || v == '+' {
			if i != 0 {
				return 0
			} else if i == 0 && v == '-' {
				sign = -1
			}
			count++
		} else if v < '0' || v > '9' || count > 1 {
			return 0
		}
		b := 0
		for k := '1'; k <= v; k++ {
			b++
		}
		a = a*10 + b
	}
	return sign * a
}
