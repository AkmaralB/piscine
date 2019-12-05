package piscine

func Split(str, charset string) []string {
	lenStr := 0
	lenCharset := 0
	word := 0
	for range str {
		lenStr++
	}
	for range charset {
		lenCharset++
	}
	for i := 0; i < lenStr-lenCharset; i++ {
		if str[i:i+lenCharset] == charset {
			word++
		}
	}
	newStr := make([]string, word+1)
	j := -1
	k := 0
	for i := 0; i < lenStr-lenCharset; i++ {
		if str[i:i+lenCharset] == charset {
			j++
			newStr[j] = str[i-k : i]
			k = 0
			i = i + lenCharset - 1
		} else {
			k++
		}
	}

	q := 0
	for _, v := range newStr {
		for range v {
			q++
		}
		q += lenCharset
	}
	q -= lenCharset
	j++
	newStr[j] = str[q:]

	return newStr
}
