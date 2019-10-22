package piscine

func LastRune(s string) rune {
	MassiveRune := []rune(s)
	ln := 0
	for range MassiveRune {
		ln++
	}

	for i, value := range MassiveRune {
		if i == ln-1 {
			return value
		}
	}
	return 0
}
