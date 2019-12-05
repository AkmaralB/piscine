package piscine

func FindNextPrime(nb int) int {
	for index := nb; index >= nb; index++ {
		if IsPrime1(index) == true {
			return index
		}
	}
	return 0
}

func IsPrime1(index int) bool {
	if index >= 2 {
		if index%2 == 0 && index != 2 {
			return false
		}
		for i := 3; i*i <= index; i = i + 2 {
			if index%i == 0 {
				return false
			}
		}
		return true
	}
	return false
}
