package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	ListArgs := os.Args

	len := 0
	for range ListArgs {
		len++
	}

	for i := len - 1; i > 0; i++ {
		for _, value := range ListArgs[i] {
			z01.PrintRune(value)
		}
	}
}
