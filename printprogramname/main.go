package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {

	AnyNameArg := os.Args
	// os.Args - massive of strings
	for _, value := range AnyNameArg[0] {
		z01.PrintRune(value)
	}
	z01.PrintRune('\n')
}
