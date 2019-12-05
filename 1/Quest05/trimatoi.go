package piscine

func TrimAtoi(s string) int {
	index := 0
	sign := 1
	for _, value := range s {
		if value >= '0' && value <= '9' {
			value = value - '0'
			index = index*10 + int(value)
		} else if index == 0 && value == '-' {
			sign = -1
		}
	}
	return sign * index
}
