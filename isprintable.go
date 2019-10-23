package piscine

func IsPrintable(str string) bool {
	RuneStr := []rune(str)
	lenStr := 0
	for range str {
		lenStr++
	}

	for i := 0; i < lenStr; i++ {
		if RuneStr[i] < ' ' {
			return false
		}
	}
	return true
}
