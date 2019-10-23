package piscine

//func len - allowed

func ToUpper(s string) string {
	RuneS := []rune(s)

	// create answer (type string)
	answer := ""

	for i := 0; i < len(RuneS); i++ {
		// print big letter for each
		if RuneS[i] >= 'a' && RuneS[i] <= 'z' {
			RuneS[i] = RuneS[i] - 32
		}
		// with each iteration send big letter to answer and transform from rune back to string type
		answer += string(RuneS[i])
	}
	return answer
}
