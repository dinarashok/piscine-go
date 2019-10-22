package piscine

func NRune(s string, n int) rune {
	MassiveRune := []rune(s)
	var x int
	x = &n
	return MassiveRune[*x]
}
