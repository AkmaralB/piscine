package piscine

func ActiveBits(n int) uint {
	count := 0
	var str string
	for n > 0 {
		str = str + string(n%2)
		n = n / 2
	}
	for _, v := range str {
		if v == 1 {
			count++
		}
	}
	return uint(count)
}
