package piscine

func SplitWhiteSpaces(str string) []string {
	var MassiveStr []string
	var b rune
	c := 0
	for i := range str {
		if i == i {
			c++
		}
	}
	for i, char := range str {
		for str[i] != ' ' {
			b = b + char
		}
		d := string(b)
		MassiveStr = append(MassiveStr, d)
	}

	return MassiveStr
}
