package piscine

func IsAlpha(str string) bool {
	RuneStr := []rune(str)
	y := 0
	for range RuneStr {
		y++
	}
	for i := 0; i < y; i++ {
		if RuneStr[i] < '0' || RuneStr[i] > '9' && RuneStr[i] < 'A' || RuneStr[i] > 'Z' && RuneStr[i] < 'a' || RuneStr[i] > 'z' {
			return false

		}
	}
	return true

}
