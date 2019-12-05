package piscine

func Unmatch(arr []int) int {
	count := 0
	match := false
	for range arr {
		count++
	}
	for i := 0; i < count; i++ {
		for j := i; j < count; j++ {
			if arr[i] == arr[j] {
				arr[j] = -1
				match = !match
			}
		}
		if match == false && arr[i] != -1 {
			return arr[i]
		}
		match = false
	}
	return -1
}
