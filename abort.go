package piscine

func Abort(a, b, c, d, e int) int {
	lens := 0
	number := []int{a, b, c, d, e}
	for range number {
		lens++
	}
	for a := 0; a < lens; a++ {
		for b := 0; b < lens; b++ {
			if number[b] < number[a] {
				number[b], number[a] = number[a], number[b]
			}
		}
	}
	return number[2]
}
