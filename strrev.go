package piscine

func StrRev(s string) string {
	initial := []rune(s)
	reverse := []rune(s)

	len := 0
	for range s {
		len++
	}

	for i = 0; i < len; i++ {
		reverse[len-i-1] = initial[i]
	}

	return string(reverse)
}
