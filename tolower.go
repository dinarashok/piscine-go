package piscine

// function len is allowed

func ToLower(s string) string {
	RuneS := []rune(s)

	for i := range RuneS {
		if RuneS[i] >= 'A' && RuneS[i] <= 'Z' {
			RuneS[i] = RuneS[i] + 32
		}
	}
	return string(RuneS)
}
