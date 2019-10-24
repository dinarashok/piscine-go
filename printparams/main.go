package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {

	// call library for list of strings Args
	ListArgs := os.Args

	//find lenth
	len := 0
	for range ListArgs {
		len++
	}
	// precise that need just parameters
	for i := 1; i <= len; i++ {
		// for range to get value
		for _, value := range ListArgs[i] {
			z01.PrintRune(value)
			z01.PrintRune('\n')
		}
	}
}
