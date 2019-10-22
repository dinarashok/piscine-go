package piscine

func TrimAtoi(s string) int {
	bool1 := false // not negative
	MassiveRune := []rune(s)
	n := 0 // lenght
	for range MassiveRune {
		n++
	}
	if n != 0 && MassiveRune[0] == '-' { //if real negative, set negative
		bool1 = true
	}
	x := 0
	for i := 0; i < n; i++ {
		if MassiveRune[i] >= '0' && MassiveRune[i] <= '9' {
			x *= 10
			x += intfor(MassiveRune[i])
		} else if MassiveRune[i] == '-' && x == 0 {
			bool1 = true
		}
	}
	if bool1 == true {
		x = x * -1
	}
	return x
}
