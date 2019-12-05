package piscine

func Map(f func(int) bool, arr []int) []bool {
	lens := 0
	for range arr {
		lens++
	}
	index := 0
	array := make([]bool, lens)
	for _, v := range arr {
		array[index] = f(v)
		index++
	}
	return array
}
