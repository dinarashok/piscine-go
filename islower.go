package piscine

func IsLower(str string) bool {
	RuneStr := []rune(str)

	y := 0
	for range str {
		y++
	}

	for i := 0; i < y; i++ {
		if RuneStr[i] < 'a' || RuneStr[i] > 'z' {
			return false
		}
	}
	return true
}
