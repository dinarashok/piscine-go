package main

import "github.com/01-edu/z01"

func main() {
	var letters = 'a'
	for letters <= 'z' {
		z01.PrintRune(rune(letters))
		letters++
	}
	z01.PrintRune('\n')
}
