package piscine

func NRune(s string, n int) rune {
	MassiveRune := []rune(s)
	for i, value := range MassiveRune {
		if i == n-1 {
			return value
		}
	}
}
