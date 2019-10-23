package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {

	AnyNameArg := []rune(os.Args[i])
	// []rune(os.Args[0]) - massive of strings
	// for _, i := range AnyNameArg {
	z01.PrintRune(rune(AnyNameArg[0]))
	// }
	z01.PrintRune('\n')
}
