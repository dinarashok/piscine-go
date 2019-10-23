package piscine

func IsNumeric(str string) bool {
	RuneStr := []rune(str)
	y := 0
	for range str {
		y++
	}

	for i := 0; i < y; i++ {
		if RuneStr[i] < '0' && RuneStr[i] > '9' {
			return false
		}
	}
	return true
}
