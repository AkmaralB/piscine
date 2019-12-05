package piscine

func Compare(a, b string) int {
	// a1 := []rune(a)
	// b1 := []rune(b)
	// sizea := 0
	// sizeb := 0
	// for range a {
	// 	sizea++
	// }
	// for range b {
	// 	sizeb++
	// }
	// for i := 0; i <= sizea; i++ {
	// 	if a1[i] > b1[i] {
	// 		return 1
	// 	} else if a1[i] < b1[i] {
	// 		return -1
	// 	} else if a1[i] == b1[i] {
	// 		continue
	// 	}
	// }
	// return 0

	if a == b {
		return 0
	} else if a > b {
		return 1
	} else {
		return -1
	}
}
