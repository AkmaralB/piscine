package piscine

func MakeRange(min, max int) []int {

	if min >= max {
		var answer []int
		return answer
	}

	answer := make([]int, max-min)
	index := 0
	for i := min; i < max; i++ {
		answer[index] = i
		index++
	}
	return answer
}
