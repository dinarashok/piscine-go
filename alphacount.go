package piscine

func AlphaCount(str string) int {
	StrInRunes := []rune(str)
	count := 0
	for i := range StrInRunes {
		if StrInRunes[i] >= 65 && StrInRunes[i] <= 132 || StrInRunes[i] >= 97 && StrInRunes[i] <= 122 {
			count++
		}
	}
	return count
}
