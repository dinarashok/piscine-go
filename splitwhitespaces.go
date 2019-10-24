package piscine

func SplitWhiteSpaces(str string) []string {
	rs := []rune(str)
	lastSep := -1
	count := 1
	for i := range rs {
		if rs[i] == '\t' || rs[i] == '\n' || rs[i] == ' ' {
			count++
		}
	}
	s := make([]string, count)
	curr := 0
	for i := range rs {
		if rs[i] == '\t' || rs[i] == '\n' || rs[i] == ' ' {
			s[curr] = string(rs[lastSep+1 : i])
			curr++
			lastSep = i
		}
	}
	s[count-1] = string(rs[lastSep+1 : my_rune_len(rs)])
	return s
}
