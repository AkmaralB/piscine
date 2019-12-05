package piscine

func SortWordArr(array []string) {
	lens := 0
	for range array {
		lens++
	}
	swapped := true
	for swapped {
		swapped = false
		for i := 1; i < lens; i++ {
			if array[i-1] > array[i] {
				array[i], array[i-1] = array[i-1], array[i]
				swapped = true
			}
		}
	}

}
