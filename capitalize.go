package piscine

func Capitalize(s string) string {
	lenS := 0
	RuneS := []rune(s)
	for range RuneS {
		lenS++
	}
	for i := 0; i < lenS; i++ {
		if RuneS[i] >= 'A' && RuneS[i] <= 'Z' {
			RuneS[i] = (RuneS[i] + 32)
		}
		if RuneS[0] >= 'a' && RuneS[0] <= 'z' {
			RuneS[0] = (RuneS[0] - 32)
		}
		if i > 0 {
			if (RuneS[i-1] > 'Z' && RuneS[i-1] < 'a') || (RuneS[i-1] < 'A' && RuneS[i-1] > '9') || RuneS[i-1] < '0' || RuneS[i-1] > 'z' {
				if RuneS[i] >= 'a' && RuneS[i] <= 'z' {
					RuneS[i] = (RuneS[i] - 32)
				}
			}
		}
	}
	answer := string(RuneS)
	return answer
}
