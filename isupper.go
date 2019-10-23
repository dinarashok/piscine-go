package piscine

func IsUpper(str string) bool {
	RuneStr := []rune(str)
	len := 0
	for range str {
		len++
	}

	for i := 0; i < len; i++ {
		if RuneS[i] < 'A' || RuneStr[i] > 'Z' {
			return false
		}
	}
	return true
}
