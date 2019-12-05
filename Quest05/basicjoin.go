package piscine

func BasicJoin(strs []string) string {
	var strNew string
	for _, v := range strs {
		strNew += v
	}
	return strNew
}
