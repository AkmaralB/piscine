package piscine

func Sqrt(nb int) int {
	var index int
	for i := 1; i <= nb; i++ {
		if i*i == nb {
			index = i
			return index
		}
	}
	return 0
}
