package piscine

func ConcatParams(test []string) string {
	var strNew string
	count := 0
	for range test {
		count++
	}
	for i, v := range test {
		if i == count-1 {
			strNew += v
		} else {
			strNew += v
			strNew += "\n"
		}
	}
	return strNew
}
