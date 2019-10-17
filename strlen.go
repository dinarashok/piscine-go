package piscine

import "github.com/01-edu/z01"

func StrLen(str string) int {
	nb := []rune(str)
	for i, value := range nb {
		z01.PrintRune(i)
	}

}
