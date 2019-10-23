package piscine

func BasicJoin(strs []string) string {
	var AnswerStr string
	for _, value := range strs {
		AnswerStr += value
	}
	return AnswerStr
}
