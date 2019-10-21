package piscine

func IterativePower(nb int, power int) int {
	x := 0
	for a := 0; a < power+1; a++ {
		x = nb * power
	}

	if power == 0 {
		return 1
	}
	return x
}
