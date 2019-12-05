package piscine

func Rot14(str string) string {
	strNew := []rune(str)
	var result string

	for i := 0; i < len(strNew); i++ {
		if strNew[i] >= 65 && strNew[i] <= 76 {
			strNew[i] = strNew[i] + 14
		} else if strNew[i] >= 77 && strNew[i] <= 90 {
			strNew[i] = strNew[i] - 12
		} else if strNew[i] >= 97 && strNew[i] <= 108 {
			strNew[i] = strNew[i] + 14
		} else if strNew[i] >= 109 && strNew[i] <= 122 {
			strNew[i] = strNew[i] - 12
		}
		result += string(strNew[i])
	}
	return result
}
