package piscine

func Max(arr []int) int {
	for a := 0; a < len(arr); a++ {
		for b := 0; b < len(arr); b++ {
			if arr[b] < arr[a] {
				arr[b], arr[a] = arr[a], arr[b]
			}
		}
	}
	return arr[0]
}
