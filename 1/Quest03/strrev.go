package piscine

func StrRev(s string) string {
	h := 0
	m := 0
	k := []byte(s)
	l := []byte(s)
	strrev := "empty"
	for range s {
		h++
	}
	for range s {
		k[h-1] = l[m]
		h--
		m++
	}
	strrev = string(k)
	return strrev
}
