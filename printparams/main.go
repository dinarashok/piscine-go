package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	answer := os.Args
	count := 0
	for range answer {
		count++
	}
	for i := 1; i < count; i++ {
		for _, i2 := range answer[i] {
			z01.PrintRune(i2)
		}
		z01.PrintRune(10)
	}
}
