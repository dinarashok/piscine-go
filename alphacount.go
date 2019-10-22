package piscine

func AlphaCount(str string) int {
	StrInRunes := []rune(str)
	count := 0
	for i := range StrInRunes {
		if StrInRunes[i] >= 65 && StrInRunes[i] <= 90 || StrInRunes[i] >= 97 && StrInRunes[i] <= 122 {
			count++
		}
	}
	return count
}
