package piscine

func Sqrt(nb int) int {
	if nb%2 == 0 {
		x := 1
		nb := nb / 2
		x += 1
	} else {
		return 0
	}
	return x
}
