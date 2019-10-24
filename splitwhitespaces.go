package piscine

func SplitWhiteSpaces(str string) []string {
	RuneStr := []rune(str)
	lastSep := -1
	count := 1
	for i := range r {
		if RuneStr[i] == '\t' || RuneStr[i] == '\n' || RuneStr[i] == ' ' {
			count++
		}
	}
	s := make([]string, count)
	curr := 0
	for i := range RuneStr {
		if RuneStr[i] == '\t' || RuneStr[i] == '\n' || RuneStr[i] == ' ' {
			s[curr] = string(RuneStr[lastSep+1 : i])
			curr++
			lastSep = i
		}
	}
	s[count-1] = RuneStr[lastSep+1 : my_rune_len(RuneStr)-1]
	return s
}
