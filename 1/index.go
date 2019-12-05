package piscine

func Index2(s string, toFind string) int {
	size1 := 0
	size2 := 0
	for range s {
		size1++
	}
	for range toFind {
		size2++
	}
	if size2 == 0 {
		return 0
	} else if size2 > size1 {
		return -1
	} else if size1 == size2 {
		if s == toFind {
			return 0
		}
		return -1
	} else if size1 > size2 {
		for i := 0; i <= size1-size2; i++ {
			if s[i:i+size2] == toFind {
				return i
				break
			}
		}
	}
	return -1
}
