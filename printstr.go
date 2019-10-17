package piscine

import "github.com/01-edu/z01"

func PrintStr(str string) {
	charOneByOne := []rune(str)
	for _, letter := range charOneByOne {
		z01.PrintRune(letter)
	}
}
