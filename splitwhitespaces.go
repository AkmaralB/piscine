package piscine

func SplitWhiteSpaces(str string) []string {
	str = clearWord(str)
	len := getWordLen(str)
	arr := make([]string, len)
	tmp := ""
	j := 0

	if !spaceCheck(str[0]) {
		tmp += string(str[0])
	}

	for i := 1; i < strLen(str); i++ {
		if spaceCheck(str[i-1]) {
			arr[j] = tmp
			tmp = string(str[i])
			j++
		} else if !spaceCheck(str[i]) {
			tmp += string(str[i])
		}
	}

	arr[j] = tmp

	return arr
}

func clearWord(s string) string {
	result := ""
	for i, v := range s {
		if i == 0 {
			if !spaceCheck(byte(v)) {
				result += string(v)
			}
		} else {
			if !(spaceCheck(s[i-1]) && spaceCheck(s[i])) {
				result += string(v)
			}
		}
	}
	len := strLen(result)
	if spaceCheck(result[len-1]) {
		result = result[:len-1]
	}
	return result
}

func spaceCheck(s byte) bool {
	return (s == ' ' || s == '\n' || s == '\t')
}

func getWordLen(s string) int {
	if strLen(s) == 0 {
		return 0
	}
	count := 1
	for _, v := range s {
		if spaceCheck(byte(v)) {
			count++

		}
	}
	return count
}

func strLen(s string) int {
	count := 0
	for range s {
		count++
	}
	return count
}
