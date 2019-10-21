package piscine

func RecursiveFactorial(nb int) int {
	if nb == 0 {
		return 1
	}
	if nb >= 1 && nb < 12 {
		x := nb * RecursiveFactorial(nb-1)
		return x
	}
	return 0
}
