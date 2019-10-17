package piscine

import "github.com/01-edu/z01.PrintRune"

func PrintStr(str string) {
	charOneByOne := []byte(str)
	for _, letter := range charOneByOne {
		z01.PrintRune("%v", letter)
	}
}
