package piscine

func Compact(ptr *[]string) int {
	lens := 0
	for _, v := range *ptr {
		if v != "" {
			lens++
		}
	}
	array := make([]string, lens)
	index := 0
	for _, v := range *ptr {
		if v != "" {
			array[index] = v
			index++
		}
	}
	*ptr = array
	return lens
}
