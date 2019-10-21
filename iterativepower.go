package piscine

func IterativePower(nb int, power int) int {
	x := 1
	if power > 0 {
		for a := 1; a < (power + 1); a++ {
			x = x * nb
		}
	}

	if power == 0 {
		return 1
	} else {
		x = 0
	}
	return x
}
