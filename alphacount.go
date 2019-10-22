package piscine

import "github.com/01-edu/z01"

func AlphaCount(str string) int {
	StrInRunes := []rune(str)
	count := 0
	for i := range StrInRunes {
		if StrInRunes[i] >= 65 && StrInrRunes[i] <= 132 || StrInRunes[i] >= 97 && StrInrRunes[i] <= 122 {
			count++
		}
	}
	return count
}
