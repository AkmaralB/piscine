package piscine

func SortIntegerTable(table []int) {
	count := 0
	for range table {
		count++
	}
	for a := 0; a < count-1; a++ {
		for b := 0; b < (count - a - 1); b++ {
			if table[b+1] < table[b] {
				table[b+1], table[b] = table[b], table[b+1]
			}
		}
	}
}
