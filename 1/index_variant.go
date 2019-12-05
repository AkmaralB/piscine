package piscine

func Index(s string, toFind string) int {
	rune1 := []rune(s)
	rune2 := []rune(toFind)
	size1 := 0
	size2 := 0
	for range s {
		size1++
	}
	for range toFind {
		size2++
	}
	if size2 == 0 {
		return -1
	}
	if size1 >= size2 {
		for i1 := 0; i1 < size1-size2; i1++ {
			if rune1[i1] == rune2[0] {
				number := 1
				for i2 := 0; i2 < size2; i2++ {
					if rune1[i1+i2] == rune2[i2] {
						continue
					} else {
						number = 0
						break
					}
				}
				if number == 1 {
					return i1
				}
			}
		}
	}
	return -1
}
