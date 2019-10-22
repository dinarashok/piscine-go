package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	var DivisionLeft []int

	if n == 0 {
		DivisionLeft = append(DivisionLeft, n)
	}
	indexNb := 0
	for n > 0 {
		DivisionLeft = append(DivisionLeft, n%10)
		n = n / 10
		indexNb++
	}
	for i := 0; i < indexNb; i++ {
		for j := i + 1; j < indexNb; j++ {
			if DivisionLeft[i] > DivisionLeft[j] {
				t := DivisionLeft[i]
				DivisionLeft[i] = DivisionLeft[j]
				DivisionLeft[j] = t
			}
		}
	}
	for i := range DivisionLeft {
		z01.PrintRune(rune(48 + DivisionLeft[i]))
	}
}
