package piscine

func RecursivePower(nb int, power int) int {
	x := nb
	if power > 0 {
		x = x * RecursivePower(nb, power-1)
	} else if power == 0 {
		x = 1
	} else {
		x = 0
	}
	return x
}
