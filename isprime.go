package piscine

func IsPrime(nb int) bool {
	prime := true
	if nb > 1 && nb < 100 {
		for a := 2; a < nb; a++ {
			if nb%a == 0 {
				prime = false
			} else {
				prime = true
			}
		}
	}
	return prime
}
