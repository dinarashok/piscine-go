package piscine

func IterativeFactorial(nb int) int {
	x := 1
	if nb > 0 && nb < 18 {
		for a := 1; a < nb+1; a++ {
			x = x * a
		}
		return x
		if nb == 0 {
			return 1
		}
	}
	return 0
}
