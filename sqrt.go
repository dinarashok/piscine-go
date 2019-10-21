package piscine

func Sqrt(nb int) int {
	if nb%2 == 0 {
		x := 1
		nb := nb / 2
		x += 1
		return x
	} else {
		return 0
	}
}
