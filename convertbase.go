package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	l := 0
	for range baseFrom {
		l++
	}
	s := ""
	//var a [100]rune
	if l < 2 { // proverka dlini massiva
		return s
	}
	r := []rune(baseFrom)
	for _, value := range baseFrom { // proverka na '+' i '-'
		if value == '-' || value == '+' {
			return s
		}
	}
	for i := 0; i < l-1; i++ { // proverka na sovpadenie
		for j := i + 1; j < l; j++ {
			if r[i] == r[j] {
				return s
			}
		}
	}

	ra := []rune(nbr)
	lr := 0
	for range ra {
		lr++
	}
	n := 0
	for i := 0; i < lr; i++ {
		count := 0

		for _, value := range r {
			if ra[i] == value {
				n = n*l + count
			}
			count++
		}
	}
	//----------------------------------------------------------------------------

	ll := 0
	for range baseTo {
		ll++
	}
	//fmt.Print("dlina =", ll, "  ")
	var a [100]rune
	if ll < 2 { // proverka dlini massiva
		return s
	}
	rr := []rune(baseTo)
	for _, value := range rr { // proverka na '+' i '-'
		if value == '-' || value == '+' {
			return s
		}
	}
	for i := 0; i < ll-1; i++ { // proverka na sovpadenie
		for j := i + 1; j < ll; j++ {
			if rr[i] == rr[j] {
				return s
			}
		}
	}

	if n < 0 { // esli otricatelnoe chislo
		return s
	}
	n1 := n
	count := 0
	for j := 1; j > 0; j++ {
		n1 = n1 / ll
		count++
		if n1 == 0 {
			break
		}
	}
	n2 := n
	for j := 0; j < count; j++ {
		//mt.Print(n2 % ll)
		a[count-1-j] = rr[n2%ll]
		n2 = n2 / ll
	}
	for j := 0; j < count; j++ {
		s += string(a[j])
	}

	return s
}
